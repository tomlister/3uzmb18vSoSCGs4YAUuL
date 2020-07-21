package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

func piece(imd *imdraw.IMDraw, pos int, symbol string) {
	imd.Color = colornames.Crimson
	imd.EndShape = imdraw.SharpEndShape
	if (pos == 0) {
		if (symbol == "x") {
			imd.Push(pixel.V(50, 768.0-50), pixel.V(256-50, 512.0+50))
			imd.Line(30)
			imd.Push(pixel.V(50, 512.0+50), pixel.V(256-50, 768.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(128, 768.0-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 1) {
		if (symbol == "x") {
			imd.Push(pixel.V(256+50, 768.0-50), pixel.V(512-50, 512.0+50))
			imd.Line(30)
			imd.Push(pixel.V(256+50, 512.0+50), pixel.V(512-50, 768.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(256+128, 768.0-128))
			imd.Circle(156 / 2 , 15)
		}
	} else if (pos == 2) {
		if (symbol == "x") {
			imd.Push(pixel.V(512+50, 768.0-50), pixel.V(768-50, 512.0+50))
			imd.Line(30)
			imd.Push(pixel.V(512+50, 512.0+50), pixel.V(768-50, 768.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(512+128, 768.0-128))
			imd.Circle(156 / 2 , 15)
		}
	} else if (pos == 3) {
		if (symbol == "x") {
			imd.Push(pixel.V(50, 512.0-50), pixel.V(256-50, 256.0+50))
			imd.Line(30)
			imd.Push(pixel.V(50, 256.0+50), pixel.V(256-50, 512.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(128, 512-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 4) {
		if (symbol == "x") {
			imd.Push(pixel.V(256+50, 512.0-50), pixel.V(512-50, 256.0+50))
			imd.Line(30)
			imd.Push(pixel.V(256+50, 256.0+50), pixel.V(512-50, 512.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(256+128, 512-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 5) {
		if (symbol == "x") {
			imd.Push(pixel.V(512+50, 512.0-50), pixel.V(768-50, 256.0+50))
			imd.Line(30)
			imd.Push(pixel.V(512+50, 256.0+50), pixel.V(768-50, 512.0-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(512+128, 512-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 6) {
		if (symbol == "x") {
			imd.Push(pixel.V(50, 256-50), pixel.V(256-50, 50))
			imd.Line(30)
			imd.Push(pixel.V(50, 50), pixel.V(256-50, 256-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(128, 256-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 7) {
		if (symbol == "x") {
			imd.Push(pixel.V(256+50, 256-50), pixel.V(512-50, 50))
			imd.Line(30)
			imd.Push(pixel.V(256+50, 50), pixel.V(512-50, 256-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(256+128, 256-128))
			imd.Circle(156 /2 , 15)
		}
	} else if (pos == 8) {
		if (symbol == "x") {
			imd.Push(pixel.V(512+50, 256-50), pixel.V(768-50, 50))
			imd.Line(30)
			imd.Push(pixel.V(512+50, 50), pixel.V(768-50, 256-50))
			imd.Line(30)
		} else if (symbol == "o") {
			imd.Push(pixel.V(512+128, 256-128))
			imd.Circle(156 /2 , 15)
		}
	}
}