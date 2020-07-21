package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type button struct {
	color         color.RGBA
	bind          *mouseareabind
	text          string
	fontsize      float64
	paddingleft   int
	paddingbottom int
}

func NewButton(color color.RGBA, bind *mouseareabind, text string, fontsize float64, paddingleft, paddingbottom int) *button {
	b := new(button)
	b.color = color
	b.bind = bind
	b.text = text
	b.fontsize = fontsize
	b.paddingleft = paddingleft
	b.paddingbottom = paddingbottom
	return b
}

func (b *button) DrawButton(win *pixelgl.Window, imd *imdraw.IMDraw, atlas *text.Atlas) {
	imd.Color = b.color
	imd.Push(pixel.V(float64(b.bind.x), float64(b.bind.y)))
	imd.Color = b.color
	imd.Push(pixel.V(float64(b.bind.x+b.bind.w), float64(b.bind.y)))
	imd.Color = b.color
	imd.Push(pixel.V(float64(b.bind.x), float64(b.bind.y+b.bind.h)))
	imd.Color = b.color
	imd.Push(pixel.V(float64(b.bind.x+b.bind.w), float64(b.bind.y+b.bind.h)))
	imd.Rectangle(4)
	btxt := text.New(pixel.V(float64(b.bind.x+b.paddingleft), float64(b.bind.y+b.paddingbottom)), atlas)
	btxt.Color = b.color
	fmt.Fprintln(btxt, b.text)
	btxt.Draw(win, pixel.IM.Scaled(btxt.Orig, b.fontsize))
}
