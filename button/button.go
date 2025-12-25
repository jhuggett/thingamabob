package button

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jhuggett/thingamabob/doodad"
	"github.com/jhuggett/thingamabob/label"
	"github.com/jhuggett/thingamabob/reaction"
)

type ButtonState int

const (
	ButtonStateNormal ButtonState = iota
	ButtonStateHovered
	ButtonStatePressed
)

type Config struct {
	OnClick func(*Button)
	label.Config
}

func New(config Config) *Button {
	button := &Button{}
	button.message = config.Message
	button.OnClick = config.OnClick

	button.labelConfig = config.Config
	button.Box = config.Layout

	if button.labelConfig.BackgroundColor == nil {
		button.labelConfig.BackgroundColor = color.RGBA{50, 50, 50, 100}
	}
	if button.labelConfig.Padding == (label.Padding{}) {
		button.labelConfig.Padding = label.Padding{Top: 5, Right: 10, Bottom: 5, Left: 10}
	}

	button.ConfigureLabelsForStates()

	button.buttonState = ButtonStateNormal

	return button
}

func (w *Button) ConfigureLabelsForStates() {
	w.labelForState = make(map[ButtonState]func() *label.Label)

	w.labelForState[ButtonStateNormal] = func() *label.Label {
		return label.New(label.Config{
			BackgroundColor: w.labelConfig.BackgroundColor,
			ForegroundColor: w.labelConfig.ForegroundColor,
			Message:         w.message,
			FontSize:        w.labelConfig.FontSize,
			Padding:         w.labelConfig.Padding,
		})
	}

	w.labelForState[ButtonStateHovered] = func() *label.Label {
		return label.New(label.Config{
			BackgroundColor: color.RGBA{50, 50, 50, 25},
			ForegroundColor: color.RGBA{
				R: 230,
				G: 255,
				B: 240,
				A: 255,
			},
			Message:  w.message,
			FontSize: w.labelConfig.FontSize,
			Padding:  w.labelConfig.Padding,
		})
	}

	w.labelForState[ButtonStatePressed] = func() *label.Label {
		return label.New(label.Config{
			BackgroundColor: color.RGBA{25, 25, 25, 50},
			ForegroundColor: color.RGBA{
				R: 130,
				G: 155,
				B: 140,
				A: 155,
			},
			Message:  w.message,
			FontSize: w.labelConfig.FontSize,
			Padding:  w.labelConfig.Padding,
		})
	}
}

type Button struct {
	message string

	labelConfig label.Config

	OnClick func(*Button)

	doodad.Default

	buttonState   ButtonState
	labelForState map[ButtonState]func() *label.Label
}

func (w *Button) Setup() {
	buttonLabel := w.labelForState[w.buttonState]()
	w.AddChild(buttonLabel)

	w.Children().Setup()

	// w.Box.Computed(func(b *box.Box) {
	// 	boundingBox := box.Bounding(w.Children().Boxes())
	// 	b.CopyDimensionsOf(boundingBox)
	// })

	w.ShrinkToFitContents()

	w.Reactions().Add(
		reaction.NewMouseUpReaction(
			doodad.MouseIsWithin[*reaction.MouseUpEvent](w),
			func(event *reaction.MouseUpEvent) {
				if w.buttonState != ButtonStatePressed {
					return
				}
				w.buttonState = ButtonStateHovered
				doodad.ReSetup(w)
				w.OnClick(w)
				event.StopPropagation()

			},
		),
		reaction.NewMouseDownReaction(
			doodad.MouseIsWithin[*reaction.MouseDownEvent](w),
			func(event *reaction.MouseDownEvent) {
				if w.buttonState == ButtonStatePressed {
					return
				}
				w.buttonState = ButtonStatePressed
				doodad.ReSetup(w)
				event.StopPropagation()
			},
		),
		reaction.NewMouseMovedReaction(
			doodad.MouseIsWithin[*reaction.MouseMovedEvent](w),
			func(event *reaction.MouseMovedEvent) {
				slog.Info("Mouse moved within button", "button", w.DebugString())
				event.StopPropagation()
				slog.Info("Mouse is within button")
				if w.buttonState == ButtonStateHovered {
					slog.Info("Button state is already Hovered")
					return
				}

				if w.buttonState == ButtonStatePressed {
					slog.Info("Button state is already Pressed")
					return
				}

				slog.Info("Setting button state to Hovered")

				w.buttonState = ButtonStateHovered
				doodad.ReSetup(w)
				ebiten.SetCursorShape(ebiten.CursorShapePointer)
			},
		),
		reaction.NewMouseMovedReaction(
			doodad.MouseIsOutside[*reaction.MouseMovedEvent](w),
			func(event *reaction.MouseMovedEvent) {
				if w.buttonState == ButtonStateNormal {
					return
				}

				w.buttonState = ButtonStateNormal
				doodad.ReSetup(w)
				ebiten.SetCursorShape(ebiten.CursorShapeDefault)
			},
		),
	)
}

func (w *Button) SetMessage(message string) {
	w.message = message
	doodad.ReSetup(w)
}

var buttonStateToDebugString = map[ButtonState]string{
	ButtonStateNormal:  "Normal",
	ButtonStateHovered: "Hovered",
	ButtonStatePressed: "Pressed",
}

func (w *Button) DebugString() string {
	return "Button: " + w.message + ", State: " + buttonStateToDebugString[w.buttonState] + ", Address: " + fmt.Sprintf("%p", w)
}

func (w *Button) DebugName() string {
	return fmt.Sprintf("Button(\"%s\")", w.message)
}
