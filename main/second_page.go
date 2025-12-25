package main

import (
	"image/color"

	"github.com/jhuggett/thingamabob/app"
	"github.com/jhuggett/thingamabob/doodad"
	"github.com/jhuggett/thingamabob/label"
	"github.com/jhuggett/thingamabob/position/box"
	"github.com/jhuggett/thingamabob/stack"
)

func NewSecondPage(
	app *app.App,
) *SecondPage {
	page := &SecondPage{
		Default: doodad.Default{},
		App:     app,
	}

	return page
}

type SecondPage struct {
	doodad.Default

	App *app.App
}

func (p *SecondPage) Setup() {
	nav := NewNavBar(p.App)
	p.AddChild(nav)

	contentStack := stack.New(stack.Config{
		BackgroundColor: color.RGBA{100, 200, 120, 100},
	})
	p.AddChild(contentStack)

	contentStack.AddChild(label.New(label.Config{
		Message: "This is the Second page",
	}))

	contentPane := box.Computed(func(b *box.Box) {
		b.Copy(p.Box).DecreaseWidth(nav.Box.Width()).MoveRight(nav.Box.Width())
	})

	contentStack.Layout().Computed(func(b *box.Box) {
		b.AlignTopWithin(contentPane).AlignLeftWithin(contentPane)
	})

	p.Children().Setup()

}
