package main

import (
	"github.com/faiface/pixel/imdraw"
)

func check(game *map[int]string, turn *int, imd *imdraw.IMDraw, tie *bool, won *bool, won2 *bool) {
	occupied := 0
	for i := 0; i < 8; i++ {
		if (*game)[i] != "" {
			occupied += 1
		}
	}
	if occupied == 8 && ((*won) == false && (*won2) == false) {
		(*tie) = true
	}
	//horizontal
	for i := 0; i < 9; i += 3 {
		if (*game)[i] == "x" && (*game)[i+1] == "x" && (*game)[i+2] == "x" {
			(*won) = true
		} else if (*game)[i] == "o" && (*game)[i+1] == "o" && (*game)[i+2] == "o" {
			(*won2) = true
		}
	}
	//vertical
	for i := 0; i < 3; i++ {
		if (*game)[i] == "x" && (*game)[i+6] == "x" && (*game)[i+3] == "x" {
			(*won) = true
		} else if (*game)[i+6] == "o" && (*game)[i+3] == "o" && (*game)[i] == "o" {
			(*won2) = true
		}
	}
	//diagonal
	if (*game)[2] == "x" && (*game)[4] == "x" && (*game)[6] == "x" {
		(*won) = true
	} else if (*game)[0] == "x" && (*game)[4] == "x" && (*game)[8] == "x" {
		(*won) = true
	} else if (*game)[2] == "o" && (*game)[4] == "o" && (*game)[6] == "o" {
		(*won2) = true
	} else if (*game)[0] == "o" && (*game)[4] == "o" && (*game)[8] == "o" {
		(*won2) = true
	}
}
