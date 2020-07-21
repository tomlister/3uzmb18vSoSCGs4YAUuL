package main

import (
	"github.com/faiface/pixel"
	"image/color"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

func shadow(imd *imdraw.IMDraw, x int, y int) {
	imd.Color = color.RGBA{0, 0, 0, 0x80}
	imd.EndShape = imdraw.NoEndShape
	imd.Push(pixel.V(float64(x+15), float64(y-15)), pixel.V(float64(x+256), float64(y-15)))
	imd.Line(30)
	imd.Push(pixel.V(float64(x+15), float64(y-256)+7.5), pixel.V(float64(x+256), float64(y-256)+7.5))
	imd.Line(30)
	imd.Color = colornames.White
	imd.EndShape = imdraw.SharpEndShape
}

func outline(imd *imdraw.IMDraw, pos int) {
	imd.Color = colornames.Crimson
	imd.EndShape = imdraw.RoundEndShape
	if (pos == 0) {
		//shadow(imd, 0, 768)
		imd.Push(pixel.V(7.5, 768.0-7.5), pixel.V(256, 512.0))
		imd.Rectangle(15)
	} else if (pos == 1) {
		//shadow(imd, 256, 768)
		imd.Push(pixel.V(256, 768.0-7.5), pixel.V(512, 512.0))
		imd.Rectangle(15)
	} else if (pos == 2) {
		//shadow(imd, 512, 768)
		imd.Push(pixel.V(512, 768.0-7.5), pixel.V(768-7.5, 512.0))
		imd.Rectangle(15)
	} else if (pos == 3) {
		//shadow(imd, 0, 512)
		imd.Push(pixel.V(7.5, 512.0), pixel.V(256, 256))
		imd.Rectangle(15)
	} else if (pos == 4) {
		imd.Push(pixel.V(256, 512), pixel.V(512, 256))
		imd.Rectangle(15)
	} else if (pos == 5) {
		imd.Push(pixel.V(512, 512), pixel.V(768-7.5, 256))
		imd.Rectangle(15)
	} else if (pos == 6) {
		imd.Push(pixel.V(7.5, 256), pixel.V(256, 7.5))
		imd.Rectangle(15)
	} else if (pos == 7) {
		imd.Push(pixel.V(256, 256), pixel.V(512, 7.5))
		imd.Rectangle(15)
	} else if (pos == 8) {
		imd.Push(pixel.V(512, 256), pixel.V(768-7.5, 7.5))
		imd.Rectangle(15)
	}
}