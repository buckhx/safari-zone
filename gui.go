package safaribot

import (
	"fmt"
	"strings"
	"time"

	"github.com/buckhx/safari-zone/util/bot"
	ui "github.com/gizak/termui"
)

const (
	enter  = "<enter>"
	space  = "<space>"
	delete = "C-8"
	ps1    = "> "
)

type GUI struct {
	bot     *SafariBot
	header  *ui.Par
	pc      ListPanel
	trainer ListPanel
	ticket  ListPanel
	display ListPanel
	input   InputPanel
}

//func NewGUI(bot *SafariBot) *GUI {
func NewGUI(opts Opts) *GUI {
	b := New(opts)
	//BotSource("bot", b.Bot)
	return &GUI{
		bot:     b,
		header:  header(),
		input:   input(),
		display: display(),
		trainer: trainer(),
		ticket:  ticket(),
		pc:      pc(),
	}
}

func (gui *GUI) Run() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(9, 0, gui.header, gui.display, gui.input),
			ui.NewCol(3, 0, gui.trainer, gui.ticket, gui.pc),
		),
	)
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Render(ui.Body)
	})
	ui.Handle("/sys/kbd", gui.input.KbdHandler)
	ui.Handle("/input/cmd", func(e ui.Event) {
		//TODO Stop on goodbye
		evt := e.Data.(InputEvt)
		gui.bot.Send(bot.Cmd(evt.Msg))
	})
	go gui.msgListen()
	ui.Body.Align()
	ui.Render(ui.Body)
	ui.Loop()
	return nil
}

func (gui *GUI) msgListen() {
	for msg := range gui.bot.Msgs {
		switch msg {
		case "":
			gui.updateTrainer()
			gui.updateTicket()
			gui.display.Clear()
		default:
			//c.display.Loading(func() {
			//	time.Sleep(1 * time.Second)
			//})
			gui.display.Append(string(msg))
		}
		ui.Render(ui.Body)
	}
	time.Sleep(1000 * time.Millisecond)
	//done channel
	ui.StopLoop()
}

func (gui *GUI) updateTicket() {
	if tkt := gui.bot.GetTicket(); tkt != nil {
		gui.ticket.BorderLabel = "TICKET"
		gui.ticket.Items = []string{
			fmt.Sprintf(" [TKTID]   %s", tkt.Uid),
			fmt.Sprintf(" [ZONE]    %s", tkt.Zone.Region),
			fmt.Sprintf(" [EXPIRES] %s", time.Duration(tkt.Expires.Time-time.Now().Unix())*time.Second),
			fmt.Sprintf(" [ENCNTRS] %d", tkt.Expires.Encounters),
			fmt.Sprintf(" [ISSUED]  %s", time.Unix(tkt.Time.Unix, 0).Format(time.RFC822)),
		}
	}
}

func (gui *GUI) updateTrainer() {
	if trn := gui.bot.GetTrainer(); trn != nil {
		gui.trainer.BorderLabel = "TRAINER"
		gui.trainer.Items = []string{
			fmt.Sprintf(" [ID]      %s", trn.Uid),
			fmt.Sprintf(" [NAME]    %s", strings.ToUpper(trn.Name)),
			fmt.Sprintf(" [GENDER]  %s", trn.Gender),
			fmt.Sprintf(" [AGE]     %d", trn.Age),
			fmt.Sprintf(" [POKEMON] %d", len(trn.Pc.Pokemon)),
		}
		gui.pc.BorderLabel = "BILL'S PC"
		gui.pc.Items = make([]string, len(trn.Pc.Pokemon))
		for i, pok := range trn.Pc.Pokemon {
			gui.pc.Items[i] = fmt.Sprintf(" [%d] %s <%s>", i, pok.NickName, pok.Name)
		}
	}
}

func trainer() ListPanel {
	pdx := ui.NewList()
	pdx.BorderLabel = "???"
	pdx.Height = 7
	pdx.Items = []string{""}
	return ListPanel{List: pdx, Format: " [%d] %s"}
}

func ticket() ListPanel {
	pdx := ui.NewList()
	pdx.BorderLabel = "???"
	pdx.Height = 7
	pdx.Items = []string{""}
	return ListPanel{List: pdx, Format: " [%d] %s"}
}

func pc() ListPanel {
	pc := ui.NewList()
	pc.BorderLabel = "???"
	pc.Height = 26
	lp := ListPanel{List: pc, Format: " [%d] %s"}
	lp.Append("Disconnected...")
	return lp
}

func header() *ui.Par {
	par := ui.NewPar(banner)
	par.Border = false
	par.Height = 8
	return par
}

func display() ListPanel {
	par := ui.NewList()
	par.BorderLabel = "Display"
	par.Height = 29
	lp := ListPanel{List: par, Prefix: " + "}
	return lp
}

func input() InputPanel {
	par := ui.NewPar(ps1)
	par.BorderLabel = "Input"
	par.Height = 3
	return InputPar("cmd", par)
}

var banner = `
   _____        __           ︵ _____
  / ____|      / _|         (⊙-)__  /
 | (___   __ _| |_ __ _ _ __ ︶  / / ___  _ __   ___
  \___ \ / _' |  _/ _' | '__|\  / / / _ \| '_ \ / _ \
  ____) | (_| | || (_| | |  | |/ /_| (_) | | | |  __/
 |_____/ \__,_|_| \__,_|_|  |_/_____\___/|_| |_|\___|`
