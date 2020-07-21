package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func logicdrawTie(mainmenu *bool, tietext *text.Text, atlas *text.Atlas, win *pixelgl.Window, imd *imdraw.IMDraw, endbutton *button, hoverstreamer *beep.StreamSeekCloser, reset func()) {
	//draw and logic for tie
	tietext.Draw(win, pixel.IM.Scaled(tietext.Orig, 6))
	endbutton.bind.CheckArea(win, func() {
		endbutton.color = colornames.Crimson
		go playhover(hoverstreamer)
	}, func() {
		(*mainmenu) = true
		reset()
	}, func() {
		endbutton.color = colornames.Black
	})
	//draw button
	endbutton.DrawButton(win, imd, atlas)
}
