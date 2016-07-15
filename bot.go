package safaribot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv/auth"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	sfrAddr = "localhost:50053"
)

type SafariBot struct {
	scanner *bufio.Scanner
	tok     string
	pdx     pbf.PokedexClient
	reg     pbf.RegistryClient
	saf     pbf.SafariClient
}

func NewSafariBot() *SafariBot {
	return &SafariBot{
		scanner: bufio.NewScanner(os.Stdin),
		reg:     registryClient(),
	}
}

func (sb *SafariBot) Run() {
	/*
		seen := sb.Greet()
		if !seen {
			sb.Register()
		}
		sb.tok = sb.SignIn()
		sb.pdx = pokedexClient(tok)
		sb.saf = safariClient(tok)
	*/
	loop(func() bool {
		_ = sb.Play()
		return !sb.yesno("Want to play again?")
	})
	/*
		say("Welcome to the Safary Zone!")
		say("Please register to participate")
		tok := register()
		say(tok)
		_, _ = pdx, saf
	*/
}

func (sb *SafariBot) Play() error {
	var rgn pbf.Zone_Code
	loop(func() bool {
		sb.say("Which region would you like to participate in? (Enter 0-6)")
		rc, err := strconv.Atoi(sb.hear())
		if err != nil || rc < 0 || rc > 6 {
			sb.say("That wasn't a valid region code. How about another?")
			return false
		}
		rgn = pbf.Zone_Code(rc)
		return sb.yesno("You'd like to visit %s?", rgn)
	})
	return nil

}

func (sb *SafariBot) Greet() (seen bool) {
	sb.say("Welcome to the Safari Zone!")
	return sb.yesno("Have you visited before?")
}

func (sb *SafariBot) Register() {
	ctx := context.Background()
	loop(func() bool {
		sb.say("What's your name?")
		name := sb.hear()
		sb.say("Hello %s, what's a secret word or phrase that we can use to identify you?", name)
		pass := sb.hear()
		var age int32
		loop(func() bool {
			sb.say("How old are you?")
			a, err := strconv.Atoi(sb.hear())
			if err != nil {
				sb.say("That's not a number! We'll have to start your registration over.")
				return false
			}
			age = int32(a)
			return true
		})
		var gdr pbf.Trainer_Gender
		loop(func() bool {
			sb.say("Are you a boy or girl? (Type boy or girl)")
			bog := sb.hear()
			switch strings.ToLower(bog) {
			case "boy":
				gdr = pbf.BOY
			case "girl":
				gdr = pbf.GIRL
			default:
				sb.say("That wasn't one of the options. Let's try again.")
				return false
			}
			return true
		})
		sb.say("Registering...")
		u, err := sb.reg.Register(ctx, &pbf.Trainer{Name: name, Password: pass, Age: age, Gender: gdr})
		if err != nil {
			sb.say("There was a problem with your registration: %q\nLet's start from the top.", err)
			return false
		}
		sb.say("Great! Here's your trainer ID %s", strings.ToLower(u.Msg))
		sb.say("You'll need to remember your trainer ID and your secret word from earlier to sign in")
		return true
	})
}

func (sb *SafariBot) SignIn() (token string) {
	ctx := context.Background()
	sb.say("Now let's sign you in to get yout token to enter different regions")
	loop(func() bool {
		sb.say("Please enter your trainer ID")
		uid := sb.hear()
		sb.say("Now your secret word please")
		pass := sb.hear()
		tok, err := sb.reg.Enter(auth.AuthenticateContext(ctx, uid, pass), &pbf.Trainer{Uid: uid, Password: pass})
		if err != nil {
			sb.say("Hmmm there was a problem %s . Let's try again", err)
			return false
		}
		token = tok.Access
		return true
	})
	return
}

/*
	saywait("Now we're getting you a token so that you can enter one of the Safari Zones", 500)
	ok = false
	for !ok {
		say("Please enter your trainer ID from above to get your token")
		uid := hear()
		tok, err := reg.Enter(auth.AuthenticateContext(ctx, uid, pass), &pbf.Trainer{Uid: uid, Password: pass})
		if err != nil {
			say("Hmmm there was a problem %s . Let's try again", err)
		} else {
			return tok.Access
		}
	}
	return ""
}
*/

func (sb *SafariBot) say(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (sb *SafariBot) hear() string {
	sb.scanner.Scan()
	return sb.scanner.Text()
}

func (sb *SafariBot) yesno(format string, v ...interface{}) (yes bool) {
	loop(func() bool {
		sb.say(format+" (Type yes or no)", v...)
		switch strings.ToLower(sb.hear()) {
		case "yes":
			yes = true
			return true
		case "no":
			yes = false
			return true
		default:
			sb.say("I didn't understand that. One more time.")
			return false
		}
	})
	return
}

func saywait(msg string, millis time.Duration) {
	fmt.Println(msg)
	time.Sleep(millis * time.Millisecond)
}

func pokedexClient(tok string) pbf.PokedexClient {
	conn, err := grpc.Dial(pdxAddr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth.AccessCredentials(tok)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pbf.NewPokedexClient(conn)
}

func registryClient() pbf.RegistryClient {
	conn, err := grpc.Dial(regAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pbf.NewRegistryClient(conn)
}

func safariClient(tok string) pbf.SafariClient {
	conn, err := grpc.Dial(sfrAddr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth.AccessCredentials(tok)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pbf.NewSafariClient(conn)
}

func loop(fn func() bool) {
	var ok bool
	for !ok {
		ok = fn()
	}
}
