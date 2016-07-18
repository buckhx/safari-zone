package safaribot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/safari"
	"github.com/buckhx/safari-zone/srv/auth"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	safAddr = "localhost:50053"
)

/*
type Opts struct {
	Safari struct {
		Addr string
	}
	Registry struct {
		Addr string
	}
}
*/

type SafariBot struct {
	reg     *registry.Client
	saf     *safari.Client
	scanner *bufio.Scanner
	ctx     context.Context
}

func New() *SafariBot {
	return &SafariBot{
		scanner: bufio.NewScanner(os.Stdin),
		ctx:     context.Background(),
	}
}

func (b *SafariBot) Connect() (err error) {
	if b.reg, err = registry.Dial(regAddr); err != nil {
		return
	}
	if b.saf, err = safari.Dial(safAddr); err != nil {
		return
	}
	return
}

func (b *SafariBot) Connected() bool {
	return b.reg != nil || b.saf != nil
}

func (b *SafariBot) Run() error {
	if !b.Connected() {
		if err := b.Connect(); err != nil {
			return err
		}
	}
	seen := b.Greet()
	if !seen {
		b.Register()
	}
	b.SignIn()
	var tkt *pbf.Ticket
	for {
		var err error
		tkt, err = b.GetTicket()
		if err != nil {
			b.say("We couldn't get you a ticket for that region %s", err)
			if !b.yes("Want to try a different region?") {
				continue
			} else {
				break
			}
		}
		break
	}
	for {
		if err := b.Encounter(tkt); err != nil {
			return err
		}
		if !b.yes("Continue walking around?") {
			break
		}
	}
	return nil
}

func (b *SafariBot) Encounter(tkt *pbf.Ticket) error {
	clms, ok := auth.ClaimsFromContext(b.ctx)
	fmt.Printf("%s - %t\n", clms, ok)
	enc, err := b.saf.Encounter(b.ctx)
	if err != nil {
		return err
	}
	if msg, err := enc.Recv(); err == nil {
		b.say(msg.Msg)
	} else {
		return err
	}
	for {
		var a *pbf.Action
		for {
			b.say("What's your move? (ball, rock, bait, run)")
			switch strings.Split(strings.ToLower(b.hear()), " ")[0] {
			case "ball":
				a = &pbf.Action{Move: &pbf.Action_Attack{"safari-ball"}}
			case "rock":
				a = &pbf.Action{Move: &pbf.Action_Attack{"throw-rock"}}
			case "bait":
				a = &pbf.Action{Move: &pbf.Action_Attack{"offer-bait"}}
			case "run":
				a = &pbf.Action{Move: &pbf.Action_Run{true}}
			case "item":
				a = &pbf.Action{Move: &pbf.Action_Item{}}
			case "switch":
				a = &pbf.Action{Move: &pbf.Action_Switch{}}
			case "":
				b.say("Gotta do something!")
				continue
			default:
				b.say("There's no time for that!")
				continue
			}
			break
		}
		if err := enc.Send(a); err != nil {
			return err
		}
		msg, err := enc.Recv()
		if err != nil {
			return err
		}
		b.say(msg.Msg)
	}
}

func (b *SafariBot) GetTicket() (tkt *pbf.Ticket, err error) {
	tid, ok := b.GetTrainerID()
	if !ok {
		err = fmt.Errorf("Unable to get trainer ID from token")
		return
	}
	for {
		b.say("Which region would you like to participate in? (Enter 0-6)")
		var zc int
		zc, err = strconv.Atoi(b.hear())
		if err != nil || zc < 0 || zc > 6 {
			b.say("That wasn't a valid region code. How about another?")
			continue
		}
		zone := &pbf.Zone{Region: pbf.Zone_Code(zc)}
		if b.yes("You'd like to visit %s?", zone.Region) {
			tkt, err = b.saf.Enter(b.ctx, &pbf.Ticket{Trainer: &pbf.Trainer{Uid: tid}, Zone: zone})
			break
		}
	}
	return
}

func (b *SafariBot) Greet() (seen bool) {
	b.say("Welcome to the Safari Zone!")
	return b.yes("Have you visited before?")
}

func (b *SafariBot) Register() {
	for {
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
			case "boy":
				gdr = pbf.BOY
			case "girl":
				gdr = pbf.GIRL
			default:
				b.say("That wasn't one of the options. Let's try again.")
				continue
			}
			break
		}
		b.say("Registering...")
		u, err := b.reg.Register(b.ctx, &pbf.Trainer{Name: name, Password: pass, Age: age, Gender: gdr})
		if err != nil {
			b.say("There was a problem with your registration: %q\nLet's start from the top.", err)
			continue
		}
		b.say("Great! Here's your trainer ID %s", strings.ToLower(u.Msg))
		b.say("You'll need to remember your trainer ID and your secret word from earlier to sign in")
		break
	}
}

func (b *SafariBot) SignIn() {
	b.say("Now let's sign you in to get your token to enter different regions")
	for {
		b.say("Please enter your trainer ID")
		uid := b.hear()
		b.say("Now your secret word please")
		pass := b.hear()
		tok, err := b.reg.Enter(auth.AuthenticateContext(b.ctx, uid, pass), &pbf.Trainer{Uid: uid, Password: pass})
		if err != nil {
			b.say("Hmmm there was a problem %s . Let's try again", err)
			continue
		}
		b.ctx = auth.AuthorizeContext(b.ctx, tok.Access)
		break
	}
	b.say("Awesome! Now we've got your token.\nYou can enter different regions for the next 24 hours before you need a new one.")
}

func (b *SafariBot) GetTrainerID() (string, bool) {
	if claims, ok := auth.ClaimsFromContext(b.ctx); ok {
		return claims.Subject, true
	}
	return "", false
}

func (b *SafariBot) say(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (b *SafariBot) hear() string {
	b.scanner.Scan()
	return b.scanner.Text()
}

func (b *SafariBot) yes(format string, v ...interface{}) (yes bool) {
	for {
		b.say(format+" (Type yes or no)", v...)
		switch strings.ToLower(b.hear()) {
		case "yes":
			return true
		case "no":
			return false
		default:
			b.say("I didn't understand that. One more time.")
		}
	}
}

func saywait(msg string, millis time.Duration) {
	fmt.Println(msg)
	time.Sleep(millis * time.Millisecond)
}
