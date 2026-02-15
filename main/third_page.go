package main

import (
	"github.com/jhuggett/thingamabob/doodad"
)

func NewThirdPage() *ThirdPage {
	page := &ThirdPage{
		Default: doodad.Default{},
	}

	return page
}

type ThirdPage struct {
	doodad.Default
}

func (p *ThirdPage) Setup() {
	nav := NewNavBar()
	p.AddChild(nav)

	p.Children().Setup()
}
