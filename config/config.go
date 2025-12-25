package config

import "image/color"

type Flow int

const (
	TopToBottom Flow = iota
	LeftToRight
)

type Padding struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

type Border struct {
	Color  color.Color
	Top    int
	Right  int
	Bottom int
	Left   int
}

func (b Border) Exists() bool {
	return b.Top > 0 || b.Right > 0 || b.Bottom > 0 || b.Left > 0
}

func EqualPadding(amount int) Padding {
	return Padding{
		Top:    amount,
		Right:  amount,
		Bottom: amount,
		Left:   amount,
	}
}

func SymmetricPadding(vertical int, horizontal int) Padding {
	return Padding{
		Top:    vertical,
		Right:  horizontal,
		Bottom: vertical,
		Left:   horizontal,
	}
}

type VerticalAlignment int

const (
	VerticalAlignmentTop VerticalAlignment = iota
	VerticalAlignmentCenter
	VerticalAlignmentBottom
)

type HorizontalAlignment int

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)
