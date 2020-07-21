package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func logicdrawMainMenu(mainmenu *bool, menutitle *text.Text, atlas *text.Atlas, win *pixelgl.Window, imd *imdraw.IMDraw, button *button, hoverstreamer *beep.StreamSeekCloser) {
	//draw and logic for mainmenu
	//draw title
	menutitle.Draw(win, pixel.IM.Scaled(menutitle.Orig, 6))
	//check if button is moused over.
	button.bind.CheckArea(win, func() {
		button.color = colornames.Crimson
		go playhover(hoverstreamer)
	}, func() {
		(*mainmenu) = false
	}, func() {
		button.color = colornames.Black
	})
	//draw button
	button.DrawButton(win, imd, atlas)
}
