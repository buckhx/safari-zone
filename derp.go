package safaribot

type Derp struct {
	*Bot
}

func (d Derp) Init() State {
	return d.Derp
}

func (d Derp) Derp() State {
	d.Msgs <- "derp"
	return d.Herp
}

func (d Derp) Herp() State {
	d.Msgs <- "herp"
	return d.Derp
}

func DerpBot() Derp {
	d := Derp{NewBot()}
	d.Run(d.Init)
	return d
}
