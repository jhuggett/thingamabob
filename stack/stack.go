package stack

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jhuggett/thingamabob/config"
	"github.com/jhuggett/thingamabob/doodad"
	"github.com/jhuggett/thingamabob/position/box"
	"github.com/jhuggett/thingamabob/reaction"

	_ "embed"
)

type LayoutRule int

const (
	FitContents LayoutRule = iota
	Fill
)

type Config struct {
	Flow         config.Flow
	SpaceBetween int
	Padding      config.Padding

	BackgroundColor color.Color

	LayoutRule LayoutRule

	HorizontalAlignment config.HorizontalAlignment
	VerticalAlignment   config.VerticalAlignment

	Border config.Border

	Shader *ebiten.Shader
}

func New(config Config) *Stack {
	s := &Stack{
		Config: config,
	}

	return s
}

type Stack struct {
	Config Config

	doodad.Default
}

/*
  - Warning: when setup is called, we reposition all children!!!
    This means that if you want to apply changes to the child of a stack, you
    should call setup first, otherwise the changes will be discarded.
*/
func (s *Stack) Setup() {
	var previousChild doodad.Doodad
	for _, child := range s.Children().Doodads {
		previousChildReferenceCopy := previousChild

		switch s.Config.Flow {
		case config.LeftToRight:
			if previousChildReferenceCopy == nil {
				child.Layout().Computed(func(b *box.Box) {
					switch s.Config.VerticalAlignment {
					case config.VerticalAlignmentTop:
						b.SetY(s.Box.Y() + s.Config.Padding.Top)
					case config.VerticalAlignmentCenter:
						b.SetY(s.Box.Y() + (s.Box.Height() / 2) - (b.Height() / 2))
					case config.VerticalAlignmentBottom:
						b.SetY(s.Box.Y() + s.Box.Height() - b.Height() - s.Config.Padding.Bottom)
					}

					b.SetX(s.Box.X() + s.Config.Padding.Left)
				})
			} else {
				child.Layout().Computed(func(b *box.Box) {
					switch s.Config.VerticalAlignment {
					case config.VerticalAlignmentTop:
						b.SetY(s.Box.Y() + s.Config.Padding.Top)
					case config.VerticalAlignmentCenter:
						b.SetY(s.Box.Y() + (s.Box.Height() / 2) - (b.Height() / 2))
					case config.VerticalAlignmentBottom:
						b.SetY(s.Box.Y() + s.Box.Height() - b.Height() - s.Config.Padding.Bottom)
					}

					b.SetX(previousChildReferenceCopy.Layout().X() +
						previousChildReferenceCopy.Layout().Width() +
						s.Config.SpaceBetween)
				})
			}
		case config.TopToBottom:
			if previousChildReferenceCopy == nil {
				child.Layout().Computed(func(b *box.Box) {
					switch s.Config.HorizontalAlignment {
					case config.HorizontalAlignmentLeft:
						b.SetX(s.Box.X() + s.Config.Padding.Left)
					case config.HorizontalAlignmentCenter:
						b.SetX(s.Box.X() + (s.Box.Width() / 2) - (b.Width() / 2))
					case config.HorizontalAlignmentRight:
						b.SetX(s.Box.X() + s.Box.Width() - b.Width() - s.Config.Padding.Right)
					}

					b.SetY(s.Box.Y() + s.Config.Padding.Top)
				})
			} else {
				child.Layout().Computed(func(b *box.Box) {
					switch s.Config.HorizontalAlignment {
					case config.HorizontalAlignmentLeft:
						b.SetX(s.Box.X() + s.Config.Padding.Left)
					case config.HorizontalAlignmentCenter:
						b.SetX(s.Box.X() + (s.Box.Width() / 2) - (b.Width() / 2))
					case config.HorizontalAlignmentRight:
						b.SetX(s.Box.X() + s.Box.Width() - b.Width() - s.Config.Padding.Right)
					}

					b.SetY(previousChildReferenceCopy.Layout().Y() +
						previousChildReferenceCopy.Layout().Height() +
						s.Config.SpaceBetween)
				})
			}
		}
		previousChild = child

	}

	s.Reactions().Add(
		reaction.NewMouseMovedReaction(
			doodad.MouseIsWithin[*reaction.MouseMovedEvent](s),
			func(event *reaction.MouseMovedEvent) {
				event.StopPropagation()
			},
		),
		reaction.NewMouseMovedReaction(
			doodad.MouseIsOutside[*reaction.MouseMovedEvent](s),
			func(event *reaction.MouseMovedEvent) {
			},
		),
		reaction.NewMouseUpReaction(
			doodad.MouseIsWithin[*reaction.MouseUpEvent](s),
			func(event *reaction.MouseUpEvent) {
				event.StopPropagation()
			},
		),
		reaction.NewMouseDragReaction(
			doodad.MouseIsWithin[*reaction.OnMouseDragEvent](s),
			func(event *reaction.OnMouseDragEvent) {
				event.StopPropagation()
			},
		),
		reaction.NewMouseWheelReaction(
			doodad.MouseIsWithin[*reaction.MouseWheelEvent](s),
			func(event *reaction.MouseWheelEvent) {
				event.StopPropagation()
			},
		),
	)
	s.Children().Setup()

	s.Box.Computed(func(b *box.Box) {
		switch s.Config.LayoutRule {
		case FitContents:

			bounding := box.Bounding(s.Children().Boxes())
			b.SetWidth(bounding.Width() + s.Config.Padding.Left + s.Config.Padding.Right)
			b.SetHeight(bounding.Height() + s.Config.Padding.Top + s.Config.Padding.Bottom)

		case Fill:
			// Do nothing, we fill the available space
		}
	})

	if (s.Config.BackgroundColor != nil || s.Config.Border.Exists()) && s.Box.Width() > 0 && s.Box.Height() > 0 {
		background := ebiten.NewImage(s.Box.Width(), s.Box.Height())
		if s.Config.BackgroundColor != nil {
			background.Fill(s.Config.BackgroundColor)
		} else {
			background.Fill(color.RGBA{0, 0, 0, 0})
		}

		// Draw the border if specified
		if s.Config.Border.Left > 0 {
			for x := 0; x < s.Config.Border.Left; x++ {
				for y := 0; y < s.Box.Height(); y++ {
					background.Set(x, y, s.Config.Border.Color)
				}
			}
		}
		if s.Config.Border.Right > 0 {
			for x := s.Box.Width() - s.Config.Border.Right; x < s.Box.Width(); x++ {
				for y := 0; y < s.Box.Height(); y++ {
					background.Set(x, y, s.Config.Border.Color)
				}
			}
		}
		if s.Config.Border.Top > 0 {
			for y := 0; y < s.Config.Border.Top; y++ {
				for x := 0; x < s.Box.Width(); x++ {
					background.Set(x, y, s.Config.Border.Color)
				}
			}
		}
		if s.Config.Border.Bottom > 0 {
			for y := s.Box.Height() - s.Config.Border.Bottom; y < s.Box.Height(); y++ {
				for x := 0; x < s.Box.Width(); x++ {
					background.Set(x, y, s.Config.Border.Color)
				}
			}
		}

		time := 0.0

		if s.Config.Shader != nil {
			s.SetCachedDraw(&doodad.CachedDraw{
				Image: background,
				Override: func(cachedDraw doodad.CachedDraw, screen *ebiten.Image) {
					x, y := s.Layout().XY()

					if s.Config.Shader != nil {

						mx, my := s.Gesturer().CurrentMouseLocation()

						mx -= x
						my -= y / 2

						time += 0.016

						opts := &ebiten.DrawRectShaderOptions{
							Uniforms: map[string]any{
								"Cursor":   []float32{float32(mx), float32(my)},
								"Radius":   float32(100),
								"Strength": float32(0.6),
							},
						}
						opts.GeoM.Translate(float64(x), float64(y))
						opts.Images[0] = background

						screen.DrawRectShader(s.Box.Width(), s.Box.Height(), s.Config.Shader, opts)
					}
				},
			})
		} else {
			s.SetCachedDraw(&doodad.CachedDraw{
				Image: background,
			})
		}
	}
}
