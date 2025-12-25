package canvas

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jhuggett/thingamabob/doodad"
)

type Canvas struct {
	doodad.Default

	Render func(screen *ebiten.Image)
}

func New(render func(screen *ebiten.Image)) *Canvas {
	return &Canvas{
		Render: render,
	}
}

func (i *Canvas) Setup() {
	// no-op
}

func (i *Canvas) Draw(screen *ebiten.Image) {
	if i.Render != nil {
		i.Render(screen)
	}
}
