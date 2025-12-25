package main

import (
	"image/color"

	"github.com/jhuggett/thingamabob/app"
	"github.com/jhuggett/thingamabob/button"
	"github.com/jhuggett/thingamabob/config"
	"github.com/jhuggett/thingamabob/doodad"
	"github.com/jhuggett/thingamabob/label"
	"github.com/jhuggett/thingamabob/position/box"
	"github.com/jhuggett/thingamabob/stack"
)

func NewNavBar(
	app *app.App,
) *NavBar {
	navBar := &NavBar{
		App: app,
	}

	return navBar
}

type NavBar struct {
	doodad.Default

	App *app.App
}

func (n *NavBar) Setup() {
	titleLabel := label.New(label.Config{
		Message:  "Design Library",
		FontSize: 24,
		ForegroundColor: color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	})

	firstPageButton := button.New(button.Config{
		OnClick: func(b *button.Button) {
			n.App.Replace(NewFirstPage(n.App))
		},
		Config: label.Config{
			Message: "First Page",
		},
	})

	secondPageButton := button.New(button.Config{
		OnClick: func(b *button.Button) {
			n.App.Replace(NewSecondPage(n.App))
		},
		Config: label.Config{
			Message: "Second Page",
		},
	})

	thirdPageButton := button.New(button.Config{
		OnClick: func(b *button.Button) {
			n.App.Replace(NewThirdPage(n.App))
		},
		Config: label.Config{
			Message: "Third Page",
		},
	})

	mainStack := stack.New(stack.Config{
		BackgroundColor: color.RGBA{
			R: 100,
			G: 150,
			B: 100,
			A: 255,
		},
		Padding:      config.SymmetricPadding(1, 2),
		SpaceBetween: 10,

		LayoutRule: stack.Fill,

		// BOXES seem to not be taking padding into account. NEED TO FIGURE THIS OUT
	})
	n.AddChild(mainStack)

	mainStack.AddChild(
		titleLabel,
		firstPageButton,
		secondPageButton,
		thirdPageButton,
	)

	n.Layout().Computed(func(b *box.Box) {
		b.SetWidth(190)
	})

	n.Children().Setup()
}
