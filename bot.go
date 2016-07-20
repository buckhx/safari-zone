package safaribot

import "fmt"

type Bottable interface {
	Init() State
	Send(Cmd)
	Get() Msg
}

type State func() State

type Msg string
type Cmd string

type Bot struct {
	Cmds chan Cmd
	Msgs chan Msg
}

func NewBot() *Bot {
	b := &Bot{
		Cmds: make(chan Cmd, 16),
		Msgs: make(chan Msg, 16),
	}
	return b
}

func (b *Bot) Send(cmd Cmd) {
	b.Cmds <- cmd
}

func (b *Bot) Get() Msg {
	return <-b.Msgs
}

func (b *Bot) Init() State {
	return b.Errorf("Init() not implemented by embedding type")
}

func (b *Bot) Errorf(format string, v ...interface{}) State {
	b.Msgs <- Msg(fmt.Sprintf(format, v...))
	return nil
}

func (b *Bot) Run(init State) {
	go b.run(init)
}

func (b *Bot) run(init State) {
	for st := init(); st != nil; {
		st = st()
	}
	close(b.Msgs)
	close(b.Cmds)
}
