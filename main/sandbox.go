package main

import (
	"image/color"

	"github.com/jhuggett/thingamabob/app"
	"github.com/jhuggett/thingamabob/button"
	"github.com/jhuggett/thingamabob/config"
	"github.com/jhuggett/thingamabob/doodad"
	"github.com/jhuggett/thingamabob/label"
	"github.com/jhuggett/thingamabob/position/box"
	"github.com/jhuggett/thingamabob/radio"
	"github.com/jhuggett/thingamabob/stack"
)

type SandboxPage struct {
	doodad.Default

	App *app.App
}

func NewSandboxPage(
	app *app.App,
) *SandboxPage {
	sandboxPage := &SandboxPage{
		App: app,
	}

	return sandboxPage
}

type YDoodad struct {
	doodad.Default
}

func (s *YDoodad) Setup() {
	stackA := stack.New(stack.Config{
		BackgroundColor: color.RGBA{100, 200, 120, 100},
	})
	s.AddChild(stackA)

	stackB := stack.New(stack.Config{
		BackgroundColor: color.RGBA{100, 200, 120, 100},
		Flow:            config.LeftToRight,
	})
	stackA.AddChild(stackB)

	stackB.AddChild(label.New(label.Config{
		Message: "Label 1",
	}))

	s.Layout().Computed(func(b *box.Box) {
		b.Copy(stackA.Box)
	})

	s.Children().Setup()
}

type XDoodad struct {
	doodad.Default
}

func (s *XDoodad) Setup() {
	stackA := stack.New(stack.Config{
		BackgroundColor: color.RGBA{200, 100, 120, 100},
		Flow:            config.LeftToRight,
	})
	s.AddChild(stackA)

	stackA.AddChild(&YDoodad{})

	s.Children().Setup()
}

func (s *SandboxPage) Setup() {
	s.AddStatefulChild("test", func() doodad.Doodad {
		return NewTestDoodad()
	})

	// xd := &XDoodad{}
	// s.AddChild(xd)

	s.Children().Setup()
}

type TestDoodad struct {
	doodad.Default
}

func NewTestDoodad() *TestDoodad {
	td := &TestDoodad{}

	return td
}

func (w *TestDoodad) Setup() {
	panel := stack.New(stack.Config{

		BackgroundColor: color.RGBA{200, 120, 20, 200},
		Padding: config.Padding{
			Top:    10,
			Right:  10,
			Bottom: 10,
			Left:   10,
		},
		SpaceBetween: 20,
	})

	w.AddChild(panel)

	timeControlRadio := doodad.Stateful(w, "radio", func() doodad.Doodad {
		return radio.New(radio.Config{
			Flow: config.LeftToRight,
			Options: []*radio.Option{
				{Label: "x0", OnSelect: func() {

				}},
				{Label: "x1", OnSelect: func() {

				}},
				{Label: "x3", OnSelect: func() {

				}},
				{Label: "x9", OnSelect: func() {

				}},
			},
			SelectedDoodad: func(option *radio.Option) doodad.Doodad {
				return button.New(button.Config{
					OnClick: func(b *button.Button) {
						option.Select()
					},
					Config: label.Config{
						Message:         option.Label,
						BackgroundColor: color.Black,
					},
				})
			},
			UnselectedDoodad: func(option *radio.Option) doodad.Doodad {
				return button.New(button.Config{
					OnClick: func(b *button.Button) {
						option.Select()
					},
					Config: label.Config{
						Message: option.Label,
					},
				})
			},
		})
	})

	// timeControlDoodad := w.Stateful("", func() doodad.Doodad {
	// 	return timeControlRadio
	// })

	panel.AddChild(timeControlRadio)

	// bottomRow := stack.New(stack.Config{
	// 	Flow: config.LeftToRight,
	// })

	// // panel.AddChild(bottomRow)

	// // currentTimeLabel := label.New(label.Config{
	// // 	Message: "Day 1 of Year 1",
	// // })

	// // bottomRow.AddChild(currentTimeLabel)

	// w.DoOnTeardown(w.Manager.OnTimeChangedCallback.Register(func(tcr outbound.TimeChangedReq) error {
	// 	w.tcr = tcr
	// 	doodad.ReSetup(w)

	// 	return nil
	// }))

	w.Children().Setup()

	// panel.Layout().Computed(func(b *box.Box) {
	// 	b.CenterHorizontallyWithin(w.Box)
	// })

	w.Layout().Recalculate()
}
