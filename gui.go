package safaribot

import ui "github.com/gizak/termui"

const (
	enter = "<enter>"
	space = "<space>"
)

func GUI() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()
	head := header()
	input := input()
	display := display()
	pdx := pokedex()
	pc := pc()
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(9, 0, head, display, input),
			ui.NewCol(3, 0, pdx, pc),
		),
	)
	ui.Handle("/sys/kbd/C-x", func(ui.Event) {
		// handle Ctrl + x combination
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd", func(e ui.Event) {
		k := e.Data.(ui.EvtKbd).KeyStr
		switch {
		//TODO DELETE
		case k == enter:
			input.Text = "> "
		case k == space:
			k = " "
			fallthrough
		case len(k) == 1:
			input.Text += k
		default:
			return
		}
		//ui.Render(input)
		ui.Render(ui.Body)
	})
	/*
		ui.Handle("/timer/1s", func(e ui.Event) {
			cnt := e.Data.(ui.EvtTimer)
			if cnt.Count%2 == 0 {
				txt := []rune(input.Text)
				txt[len(txt)-1] = ' '
				input.Text = string(txt)
			} else {
				txt := []rune(input.Text)
				txt[len(txt)-1] = '_'
				input.Text = string(txt)
			}
			ui.Render(ui.Body)
		})
	*/
	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Render(ui.Body)
	})
	ui.Body.Align()
	ui.Render(ui.Body)
	ui.Loop()
	return nil
}

func pokedex() *ui.List {
	pdx := ui.NewList()
	pdx.BorderLabel = "Pokedex"
	pdx.Height = 20
	pdx.Items = []string{"Register to use the pokedex"}
	return pdx
}

func pc() *ui.List {
	pc := ui.NewList()
	pc.BorderLabel = "Bill's PC"
	pc.Height = 20
	pc.Items = []string{"Register to access the PC"}
	return pc
}

func header() *ui.Par {
	par := ui.NewPar(banner)
	par.Border = false
	par.Height = 8
	return par
}

func input() *ui.Par {
	par := ui.NewPar("> ")
	par.BorderLabel = "Input"
	par.Height = 3
	return par
}

func display() *ui.Par {
	par := ui.NewPar("Welcome to the Safari Zone!")
	par.BorderLabel = "Display"
	par.Height = 29
	return par
}
