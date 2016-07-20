package safaribot

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui"
)

type ListPanel struct {
	*ui.List
	Format string
	Prefix string
}

func (d ListPanel) Append(lines ...string) {
	for i, line := range lines {
		switch {
		case d.Format != "":
			line = fmt.Sprintf(d.Format, i+len(d.Items), line)
		case d.Prefix != "":
			line = fmt.Sprint(d.Prefix, line)
		}
		lines[i] = line
	}
	d.Items = append(d.Items, lines...)
}

func (d ListPanel) Clear() {
	d.Items = []string{}
}

type InputPanel struct {
	*ui.Par
	name string
	evts chan<- ui.Event
}

// Name is used to identify events to the mux "/input/<name>"
func InputPar(name string, par *ui.Par) InputPanel {
	evts := make(chan ui.Event, 32)
	ui.DefaultEvtStream.Merge(name, evts)
	return InputPanel{
		Par:  par,
		name: name,
		evts: evts,
	}
}

func (in InputPanel) Backspace() bool {
	if len(in.Text) > len(ps1) {
		in.Text = in.Text[:len(in.Text)-1]
		return true
	}
	return false
}

func (in InputPanel) Clear() string {
	v := in.Value()
	in.Text = ps1
	return v
}

func (in InputPanel) Value() string {
	return in.Text[len(ps1):]
}

func (in InputPanel) Append(val string) string {
	in.Text += val
	return in.Text
}

func (in InputPanel) KbdHandler(e ui.Event) {
	k := e.Data.(ui.EvtKbd).KeyStr
	switch {
	case k == enter:
		in.emit(in.Clear())
	case k == delete:
		in.Backspace()
	case k == space:
		k = " "
		fallthrough
	case len(k) == 1:
		in.Append(k)
	default:
		return
	}
	ui.Render(in)
}

func (in InputPanel) emit(msg string) {
	in.evts <- ui.Event{
		Type: "input",
		Path: "/input/" + in.name,
		Time: time.Now().Unix(),
		Data: InputEvt{
			Msg: msg,
			Src: in.name,
		},
	}
}

type InputEvt struct {
	Msg, Src string
}
