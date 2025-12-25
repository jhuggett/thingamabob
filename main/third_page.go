package main

import (
	"github.com/jhuggett/thingamabob/app"
	"github.com/jhuggett/thingamabob/doodad"
)

func NewThirdPage(
	app *app.App,
) *ThirdPage {
	page := &ThirdPage{
		Default: doodad.Default{},
		App:     app,
	}

	return page
}

type ThirdPage struct {
	doodad.Default

	App *app.App
}

func (p *ThirdPage) Setup() {
	nav := NewNavBar(p.App)
	p.AddChild(nav)

	p.Children().Setup()
}
