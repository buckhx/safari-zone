package safaribot

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/safari"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/buckhx/safari-zone/util/bot"
)

type CtxKey int

const (
	TrainerKey CtxKey = iota
	TicketKey
)

type SafariBot struct {
	bot.Bot
	opts Opts
	reg  *registry.Client
	saf  *safari.Client
	ctx  context.Context
}

func New(opts Opts) *SafariBot {
	d := &SafariBot{
		Bot:  bot.New(),
		opts: opts,
		ctx:  context.Background(),
	}
	d.Run(d.Init)
	return d
}

func (b *SafariBot) Init() bot.State {
	return b.Connect
}

func (b *SafariBot) Connect() bot.State {
	//b.say("Connecting...")
	var err error
	if b.reg, err = b.opts.DialRegistry(); err != nil {
		return b.Errorf("Couldn't connect %q", err)
	}
	if b.saf, err = b.opts.DialSafari(); err != nil {
		return b.Errorf("Couldn't connect %q", err)
	}
	uid := os.Getenv("SAFARI_UID")
	pass := os.Getenv("SAFARI_PASS")
	if uid != "" && pass != "" {
		return b.setcreds(uid, pass)
	}
	b.say("Welcome to the Safari Zone!")
	if b.yes("Would you like to play?") {
		return b.SignIn
	}
	return b.Exit
}

func (b *SafariBot) setcreds(uid, pass string) bot.State {
	//TODO return err instead of state
	u := &pbf.Trainer{Uid: uid, Password: pass}
	ctx := auth.AuthenticateContext(b.ctx, u.Uid, u.Password)
	tok, err := b.reg.Enter(ctx, u)
	if err != nil {
		panic(err)
	}
	ctx = auth.AuthorizeContext(b.ctx, tok.Access)
	trn, err := b.reg.GetTrainer(ctx, u)
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, TrainerKey, trn)
	tkt, err := b.saf.Enter(ctx, &pbf.Ticket{Trainer: trn, Zone: &pbf.Zone{Region: pbf.KANTO}})
	if err != nil {
		panic(err)
	}
	b.ctx = context.WithValue(ctx, TicketKey, tkt)
	return b.WalkAround
}

func (b *SafariBot) SignIn() bot.State {
	if !b.yes("Are you a registered trainer?") {
		return b.Register
	}
	b.say("Let's sign you in to get your token to enter different regions")
	for {
		u := b.GetTrainer()
		if u == nil || !b.yes("Is your Trainer ID %s?", u.Uid) {
			b.say("Please enter your Trainer ID")
			uid := b.hear()
			u = &pbf.Trainer{Uid: uid}
		}
		b.say("Enter your secret please")
		pass := b.hear()
		authctx := auth.AuthenticateContext(b.ctx, u.Uid, pass)
		tok, err := b.reg.Enter(authctx, u)
		//tok, err := b.reg.Enter(auth.AuthenticateContext(b.ctx, u.Uid, u.Password), u)
		if err != nil {
			b.say("Hmmm there was a problem %q. Let's try again", grpc.ErrorDesc(err))
			continue
		}
		b.ctx = auth.AuthorizeContext(b.ctx, tok.Access)
		trn, err := b.reg.GetTrainer(b.ctx, u)
		if err != nil {
			b.say("Hmmm there was a problem %q. Let's try again", grpc.ErrorDesc(err))
			continue
		}
		b.ctx = context.WithValue(b.ctx, TrainerKey, trn)
		break
	}
	b.say("Awesome! Now we've got your token")
	b.say("You can get tickets for different regions for the next 24 hours before you need a new token.")
	return b.FetchTicket
}

func (b *SafariBot) Register() bot.State {
	b.say("What's your name?")
	name := b.hear()
	b.say("Hello %s, what's a secret word or phrase that we can use to identify you?", name)
	pass := b.hear()
	var age int32
	for {
		b.say("How old are you?")
		a, err := strconv.Atoi(b.hear())
		if err != nil {
			b.say("That's not a number! We'll have to start your registration over.")
			continue
		}
		age = int32(a)
		break
	}
	var gdr pbf.Trainer_Gender
	for {
		b.say("Are you a boy or girl? (Type boy or girl)")
		bog := b.hear()
		switch strings.ToLower(bog) {
		case "b", "boy":
			gdr = pbf.BOY
		case "g", "girl":
			gdr = pbf.GIRL
		default:
			b.say("That wasn't one of the options. Let's try again.")
			continue
		}
		break
	}
	b.say("Registering...")
	trn, err := b.reg.Register(b.ctx, &pbf.Trainer{Name: name, Password: pass, Age: age, Gender: gdr})
	if err != nil {
		b.say("There was a problem with your registration %q", grpc.ErrorDesc(err))
		if b.yes("Say ok to start from the top") {
			return b.Register
		}
		return b.Exit
	}
	b.ctx = context.WithValue(b.ctx, TrainerKey, trn)
	b.say("Great! %s", trn.Name)
	b.say("You'll need to remember your trainer ID and your secret word from earlier to sign in")
	return b.SignIn
}

func (b *SafariBot) FetchTicket() bot.State {
	trn := b.GetTrainer()
	if trn == nil {
		b.say("Couldn't verify your token. Let's get a new one.")
		return b.SignIn
	}
	b.say("Which region would you like to participate in? (Enter 0-6)")
	zc, err := strconv.Atoi(b.hear())
	if err != nil || zc < 0 || zc > 6 {
		b.say("That wasn't a valid region code. How about another?")
		return b.FetchTicket
	}
	zone := &pbf.Zone{Region: pbf.Zone_Code(zc)}
	if !b.yes("You'd like to visit %s?", zone.Region) {
		return b.FetchTicket
	}
	tkt, err := b.saf.Enter(b.ctx, &pbf.Ticket{Trainer: trn, Zone: zone})
	if err != nil {
		return b.Errorf("There was a problem geting your ticket %s", grpc.ErrorDesc(err))
	}
	b.ctx = context.WithValue(b.ctx, TicketKey, tkt)
	return b.WalkAround
}

func (b *SafariBot) WalkAround() bot.State {
	if b.yes("Walk around?") {
		if rand.Float32() <= 0.5 {
			return b.Encounter
		} else {
			b.say("What a lovely day!")
			time.Sleep(1 * time.Second)
			return b.WalkAround
		}
	}
	return b.Exit
}

func (b SafariBot) Encounter() bot.State {
	b.say("Encounter!")
	return b.WalkAround
}

func (b *SafariBot) Exit() bot.State {
	b.say("Goodbye!")
	return nil
}

func (b *SafariBot) GetTrainer() *pbf.Trainer {
	u, _ := b.ctx.Value(TrainerKey).(*pbf.Trainer)
	return u
}

func (b *SafariBot) GetTicket() *pbf.Ticket {
	u, _ := b.ctx.Value(TicketKey).(*pbf.Ticket)
	return u
}

func (b *SafariBot) Context() context.Context {
	return b.ctx
}

func (b *SafariBot) say(format string, v ...interface{}) {
	b.Msgs <- bot.Msg(fmt.Sprintf(format, v...))
}

func (b *SafariBot) hear() string {
	return string(<-b.Cmds)
}

func (b *SafariBot) yes(format string, v ...interface{}) bool {
	b.say(format, v...)
	for {
		switch strings.ToLower(b.hear()) {
		case "y", "yes", "ok", "true":
			return true
		case "n", "no", "false", "nope":
			return false
		case "what", "huh", "repeat", "again":
			b.say(format, v...)
		default:
			b.say("Please answer yes or no")
		}
	}
}
