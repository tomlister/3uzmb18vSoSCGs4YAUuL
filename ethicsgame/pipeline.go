package main

import "github.com/hajimehoshi/ebiten"

func (engine *Engine) update(screen *ebiten.Image) error {
	engine.StoryData.run(screen, engine.AssetData)
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	return nil
}
