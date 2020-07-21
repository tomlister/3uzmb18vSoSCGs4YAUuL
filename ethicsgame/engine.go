package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	WindowName string
	StoryData  *Story
	AssetData  *Assets
}

func (engine *Engine) start() {
	if err := ebiten.Run(engine.update, 320, 240, 2, engine.WindowName); err != nil {
		log.Fatal(err)
	}
}
