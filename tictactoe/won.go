package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func logicdrawWon(won, won2, mainmenu *bool, wontext, won2text *text.Text, atlas *text.Atlas, win *pixelgl.Window, imd *imdraw.IMDraw, endbutton *button, hoverstreamer *beep.StreamSeekCloser, reset func()) {
	//draw and logic for won
	if *won {
		wontext.Draw(win, pixel.IM.Scaled(wontext.Orig, 6))
	} else if *won2 {
		won2text.Draw(win, pixel.IM.Scaled(won2text.Orig, 6))
	}
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
