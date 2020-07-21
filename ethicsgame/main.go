package main

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

var (
	normalFont    font.Face
	prevMouseDown bool = false
)

func init() {
	//fonts
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main() {
	//main
	assets := Assets{}
	assets.importManifest()
	story := parseStory("story.json")
	engine := Engine{
		WindowName: "Ethics Game",
		StoryData:  story,
		AssetData:  &assets,
	}
	engine.start()
}
