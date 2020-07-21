package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type boardButtons struct {
	topLeft      *mouseareabind
	topMiddle    *mouseareabind
	topRight     *mouseareabind
	middleLeft   *mouseareabind
	middle       *mouseareabind
	middleRight  *mouseareabind
	bottomLeft   *mouseareabind
	bottomMiddle *mouseareabind
	bottomRight  *mouseareabind
}

func logicdrawBoard(tie, placed, won, won2, mainmenu *bool, board *map[int]string, win *pixelgl.Window, imd *imdraw.IMDraw, boardbtns boardButtons, starter, turn, pos *int) {
	//draw and logic for board
	//draw board
	imd.Color = colornames.Black
	imd.Push(pixel.V(0, 256), pixel.V(768, 256))
	imd.Line(15)
	imd.Push(pixel.V(0, 512), pixel.V(768, 512))
	imd.Line(15)
	imd.Push(pixel.V(256, 0), pixel.V(256, 768))
	imd.Line(15)
	imd.Push(pixel.V(512, 0), pixel.V(512, 768))
	imd.Line(15)
	//draw pieces
	for i := 0; i < 9; i++ {
		piece(imd, i, (*board)[i])
	}
	//handle mouse
	//converted legacy code to new system
	(*boardbtns.topLeft).CheckArea(win, func() {
		(*pos) = 0
	}, func() {}, func() {})
	(*boardbtns.topMiddle).CheckArea(win, func() {
		(*pos) = 1
	}, func() {}, func() {})
	(*boardbtns.topRight).CheckArea(win, func() {
		(*pos) = 2
	}, func() {}, func() {})
	(*boardbtns.middleLeft).CheckArea(win, func() {
		(*pos) = 3
	}, func() {}, func() {})
	(*boardbtns.middle).CheckArea(win, func() {
		(*pos) = 4
	}, func() {}, func() {})
	(*boardbtns.middleRight).CheckArea(win, func() {
		(*pos) = 5
	}, func() {}, func() {})
	(*boardbtns.bottomLeft).CheckArea(win, func() {
		(*pos) = 6
	}, func() {}, func() {})
	(*boardbtns.bottomMiddle).CheckArea(win, func() {
		(*pos) = 7
	}, func() {}, func() {})
	(*boardbtns.bottomRight).CheckArea(win, func() {
		(*pos) = 8
	}, func() {}, func() {})
	//Render Outline
	outline(imd, (*pos))
	//Game Board Key Press
	if win.JustPressed(pixelgl.MouseButtonLeft) && (*turn) == 0 {
		if (*turn) == 0 {
			if (*board)[(*pos)] == "" {
				(*board)[(*pos)] = "x"
				(*turn) = 1
			}
		}
	} else {
		alg(board, placed, turn, starter)
	}
	check(board, turn, imd, tie, won, won2)
}
