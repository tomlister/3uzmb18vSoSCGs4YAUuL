package main

import (
	"encoding/json"
	"errors"
	"image/color"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type StoryContentImage struct {
	Image string `json:"image"`
}

type StoryContentText struct {
	Text string `json:"text"`
}

type StoryFrameIntermediate struct {
	ID      int             `json:"id"`
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

type StoryFrame struct {
	ID      int
	Type    string
	Content interface{}
}

type Story struct {
	Location int
	Frames   map[string]StoryFrame
}

func (story *Story) getKey(ID int) (string, error) {
	for k, v := range story.Frames {
		if v.ID == ID {
			return k, nil
		}
	}
	return "", errors.New("[Story] Unable to find Key from ID")
}

func (story *Story) getID(key string) (int, error) {
	for k, v := range story.Frames {
		if k == key {
			return v.ID, nil
		}
	}
	return -1, errors.New("[Story] Unable to find ID from Key")
}

func parseStory(fname string) *Story {
	data, err := ioutil.ReadFile(fname)
	e(err)
	story := &Story{
		Location: 0,
		Frames:   make(map[string]StoryFrame),
	}
	intermediatemap := make(map[string]StoryFrameIntermediate)
	err = json.Unmarshal([]byte(data), &intermediatemap)
	e(err)
	for k, v := range intermediatemap {
		var resolved interface{}
		switch v.Type {
		case "image":
			resolved = StoryContentImage{}
		case "text":
			resolved = StoryContentText{}
		}
		err = json.Unmarshal(v.Content, &resolved)
		//fmt.Println(reflect.TypeOf(resolved))
		story.Frames[k] = StoryFrame{
			ID:      v.ID,
			Type:    v.Type,
			Content: resolved,
		}
		e(err)
	}
	return story
}

func (story *Story) renderImage(screen *ebiten.Image, AssetData *Assets, imageAssetName string) {
	imageID, err := AssetData.getID(imageAssetName)
	e(err)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0.0, 0.0)
	screen.DrawImage(AssetData.data[imageID], options)
}

func (story *Story) renderChoiceButton(screen *ebiten.Image, textData string, pos int, action float64) {
	hoverymin := 180
	hoverymax := 210
	if pos == 0 {
		hoverymin = 210
		hoverymax = 240
	}
	_, my := ebiten.CursorPosition()
	if my > hoverymin && my < hoverymax {
		background, _ := ebiten.NewImage(320, 30, ebiten.FilterDefault)
		opts := &ebiten.DrawImageOptions{}
		r := float64(0x77)
		g := float64(0x77)
		b := float64(0x77)
		opts.ColorM.Translate(r, g, b, 0.2)
		opts.GeoM.Translate(0, float64(hoverymin))
		screen.DrawImage(background, opts)
		background.Dispose()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			prevMouseDown = true
		} else {
			if prevMouseDown == true {
				story.Location = int(action)
			}
			prevMouseDown = false
		}
	}
	text.Draw(screen, textData, normalFont, 5, hoverymax-10, color.White)
}

func (story *Story) run(screen *ebiten.Image, AssetData *Assets) {
	key, err := story.getKey(story.Location)
	e(err)
	switch story.Frames[key].Type {
	case "image":
		disgustingtypeconversionimsorry := story.Frames[key].Content.(map[string]interface{})["image"].(string)
		story.renderImage(screen, AssetData, disgustingtypeconversionimsorry)
		choice0Text := story.Frames[key].Content.(map[string]interface{})["choice0_text"].(string)
		choice0Action := story.Frames[key].Content.(map[string]interface{})["choice0_action"].(float64)
		story.renderChoiceButton(screen, choice0Text, 0, choice0Action)
		choice1Text := story.Frames[key].Content.(map[string]interface{})["choice1_text"].(string)
		choice1Action := story.Frames[key].Content.(map[string]interface{})["choice1_action"].(float64)
		story.renderChoiceButton(screen, choice1Text, 1, choice1Action)
	}
}
