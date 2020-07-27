package text32x36

import (
	"github.com/reiver/go-font8x8"

	"image"
	"image/color"
)

const (
	imageWidth  = 32*8
	imageHeight =  36*8
)

// At returns the color of the pixel at (x,y).
//
// At helps text32x36.Slice fit the Go built-in image.Image interface.
func (receiver Slice) At(x, y int) color.Color {
	if nil == receiver {
		return nil
	}

	if x < 0 || imageWidth <= x {
		return nil
	}
	if y < 0 || imageHeight <= y {
		return nil
	}

	var offset int
	{
		xx := x / 8
		yy := y / 8

		offset = receiver.runesOffset(xx,yy)
	}

	r := receiver.Runes()[offset]

	var c color.Color
	{
		font := font8x8.Rune(r)

		xx := x % 8
		yy := y % 8

		c = font.At(xx,yy)
	}

	return c
}

// Bounds returns the rectangular area for which there may be non-zero colors.
//
// BOunds helps text32x36.Slice fit the Go built-in image.Image interface.
func (receiver Slice) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	return image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x+imageWidth,
			Y: y+imageHeight,
		},
	}
}

// ColorModel returns the color model.
//
// BOunds helps text32x36.Slice fit the Go built-in image.Image interface.
func (receiver Slice) ColorModel() color.Model {
	const anyRune rune = 'A'

	return font8x8.Rune(anyRune).ColorModel()
}
