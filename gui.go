package safaribot

import (
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
	info    ListPanel
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
		info:    infoPanel(),
		pc:      pc(),
	}
}

func (c *GUI) Run() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(9, 0, c.header, c.display, c.input),
			ui.NewCol(3, 0, c.pc, c.info),
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
	ui.Handle("/sys/kbd", c.input.KbdHandler)
	ui.Handle("/input/cmd", func(e ui.Event) {
		evt := e.Data.(InputEvt)
		c.bot.Send(bot.Cmd(evt.Msg))
	})
	go func() {
		for msg := range c.bot.Msgs {
			switch msg {
			case "":
				//c.display.Clear()
			default:
				c.display.Append("...")
				ui.Render(ui.Body)
				time.Sleep(500 * time.Millisecond)
				c.display.Trim()
				c.display.Append(string(msg))
			}
			//m.Unlock()
			ui.Render(ui.Body)
		}
	}()
	ui.Body.Align()
	ui.Render(ui.Body)
	ui.Loop()
	return nil
}

func infoPanel() ListPanel {
	pdx := ui.NewList()
	pdx.BorderLabel = "???"
	pdx.Height = 20
	pdx.Items = []string{""}
	return ListPanel{List: pdx, Format: " [%d] %s"}
}

func pc() ListPanel {
	pc := ui.NewList()
	pc.BorderLabel = "Bill's PC"
	pc.Height = 20
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
