package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Asset struct {
	Location string `json:"location"`
	ID       int    `json:"-"`
}

type Assets struct {
	manifest map[string]Asset
	data     []*ebiten.Image
}

func (a *Assets) getID(key string) (int, error) {
	for k, v := range a.manifest {
		if k == key {
			return v.ID, nil
		}
	}
	return -1, errors.New("[Assets] Unable to find ID from Key")
}

func loadImage(fname string) *ebiten.Image {
	image, _, _ := ebitenutil.NewImageFromFile(fname, ebiten.FilterDefault)
	return image
}

func (a *Assets) importManifest() {
	data, err := ioutil.ReadFile("assets/manifest.json")
	e(err)
	err = json.Unmarshal([]byte(data), &a.manifest)
	e(err)
	for k, v := range a.manifest {
		imageasset := loadImage("assets/" + v.Location)
		a.data = append(a.data, imageasset)
		v.ID = len(a.data) - 1
		a.manifest[k] = v
		fmt.Printf("[Assets] Loaded: %s\n", k)
	}
}
