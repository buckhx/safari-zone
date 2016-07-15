package main

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

var (
	scanner *bufio.Scanner
	pdx     pbf.PokedexClient
	reg     pbf.RegistryClient
	saf     pbf.SafariClient
)

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	reg = registryClient()
	say("Welcome to the Safary Zone!")
	say("Please register to participate")
	tok := register()
	pdx = pokedexClient(tok)
	saf = safariClient(tok)
	say(tok)
	_, _ = pdx, saf
}

func register() string {
	ctx := context.Background()
	var pass string
	var ok bool
	for !ok {
		say("What's your name?")
		name := hear()
		say("Hello %s, what's a secret word or phrase that we can use to identify you?", name)
		pass = hear()
		say("How old are you?")
		age, err := strconv.Atoi(hear())
		if err != nil {
			say("That's not a number! We'll have to start your registration over.")
			continue
		}
		say("Are you a boy or girl? (Type boy or girl)")
		bog := hear()
		var gdr pbf.Trainer_Gender
		switch strings.ToLower(bog) {
		case "boy":
			gdr = pbf.BOY
		case "girl":
			gdr = pbf.GIRL
		default:
			say("That wasn't one of the options. Let's start from the top.")
			continue
		}
		say("Registering...")
		u, err := reg.Register(ctx, &pbf.Trainer{Name: name, Password: pass, Age: int32(age), Gender: gdr})
		if err != nil {
			say("There was a problem with your registration: %q\nLet's try again.", err)
		} else {
			ok = true
			saywait("You're registered! Here's your trainer ID "+u.Msg, 500)
		}
	}
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

func saywait(msg string, millis time.Duration) {
	fmt.Println(msg)
	time.Sleep(millis * time.Millisecond)
}

func say(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func hear() string {
	scanner.Scan()
	return scanner.Text()
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
