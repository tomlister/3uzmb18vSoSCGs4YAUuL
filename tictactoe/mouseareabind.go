package main

import "github.com/faiface/pixel/pixelgl"

type mouseareabind struct {
	hovering bool
	x        int
	y        int
	w        int
	h        int
}

//NewMouseAreaBind Creates a new mouseareabind with params
func NewMouseAreaBind(x, y, width, height int) *mouseareabind {
	mab := new(mouseareabind)
	mab.x = x
	mab.y = y
	mab.w = width
	mab.h = height
	return mab
}

func offhovertest(mab *mouseareabind, offhover func()) {
	if mab.hovering == true {
		mab.hovering = false
		offhover()
	}
}

func (mab *mouseareabind) CheckArea(win *pixelgl.Window, onhover, onclick, offhover func()) {
	if int(win.MousePosition().X) >= mab.x && int(win.MousePosition().X) <= mab.x+mab.w {
		if int(win.MousePosition().Y) >= mab.y && int(win.MousePosition().Y) <= mab.y+mab.h {
			if mab.hovering == false {
				onhover()
				mab.hovering = true
			}
			if win.JustPressed(pixelgl.MouseButtonLeft) {
				onclick()
			}
		} else {
			offhovertest(mab, offhover)
		}
	} else {
		offhovertest(mab, offhover)
	}
}
