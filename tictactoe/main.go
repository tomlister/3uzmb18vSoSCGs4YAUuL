package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/beep"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func FUCKAPPLE(win *pixelgl.Window) {
	pos := win.GetPos()
	win.SetPos(pixel.ZV)
	win.SetPos(pos)
}

func playhover(hoverstreamer *beep.StreamSeekCloser) {
	speaker.Lock()
	speaker.Play(*hoverstreamer)
	speaker.Unlock()
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tic Tac Toe",
		Bounds: pixel.R(0, 0, 768, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	tietext := text.New(pixel.V(256+65, 256+100), atlas)
	tietext.Color = color.RGBA{0x14, 0x3c, 0xdb, 0xff}
	fmt.Fprintln(tietext, "Tie")
	menutitle := text.New(pixel.V(150, 500), atlas)
	menutitle.Color = color.RGBA{0x00, 0x00, 0x00, 0xff}
	fmt.Fprintln(menutitle, "Tic Tac Toe")
	wontext := text.New(pixel.V(256+30, 256+100), atlas)
	wontext.Color = color.RGBA{0x00, 0x00, 0x00, 0xff}
	fmt.Fprintln(wontext, "X Won")
	won2text := text.New(pixel.V(256+30, 256+100), atlas)
	won2text.Color = color.RGBA{0x00, 0x00, 0x00, 0xff}
	fmt.Fprintln(won2text, "O Won")
	//fix apple's opengl deprecation
	FUCKAPPLE(win)

	//load sounds
	hoverf, err := os.Open("hover.mp3")
	if err != nil {
		panic(err)
	}
	hoverstreamer, hoverformat, err := mp3.Decode(hoverf)
	if err != nil {
		log.Fatal(err)
	}
	//defer hoverstreamer.Close()
	speaker.Init(hoverformat.SampleRate, hoverformat.SampleRate.N(time.Second/10))

	board := make(map[int]string, 9)
	s1 := rand.NewSource(time.Now().UTC().UnixNano())
	r1 := rand.New(s1)
	random_num := r1.Intn(3)
	pos, turn, starter := 0, 0, 0
	if random_num == 0 {
		turn = 0
		starter = 0
	} else if random_num == 1 {
		turn = 1
		starter = 1
	} else if random_num == 2 {
		turn = 0
		starter = 0
	} else if random_num == 3 {
		turn = 1
		starter = 1
	}
	tie, won, won2, placed, mainmenu := false, false, false, false, true

	//the wrong way to define a function. nested functions in go is yuck
	reset := func() {
		tie, won, won2, placed = false, false, false, false
		board = make(map[int]string, 9)
		pos, turn, starter = 0, 0, 0
		random_num = r1.Intn(3)
		if random_num == 0 {
			turn = 0
			starter = 0
		} else if random_num == 1 {
			turn = 1
			starter = 1
		} else if random_num == 2 {
			turn = 0
			starter = 0
		} else if random_num == 3 {
			turn = 1
			starter = 1
		}
	}

	//menu buttons
	buttonBind := NewMouseAreaBind((768/2)-50, (768/2)-200, 100, 50)
	button := NewButton(colornames.Black, buttonBind, "Start", 2, 16, 15)

	endbuttonBind := NewMouseAreaBind((768/2)-50, (768/2)-200, 100, 50)
	endbutton := NewButton(colornames.Black, endbuttonBind, "Back", 2, 22, 15)

	//board mouse area binds
	topLeft := NewMouseAreaBind(0, 512, 256, 256)
	topMiddle := NewMouseAreaBind(256, 512, 256, 256)
	topRight := NewMouseAreaBind(512, 512, 256, 256)
	middleLeft := NewMouseAreaBind(0, 256, 256, 256)
	middle := NewMouseAreaBind(256, 256, 256, 256)
	middleRight := NewMouseAreaBind(512, 256, 256, 256)
	bottomLeft := NewMouseAreaBind(0, 0, 256, 256)
	bottomMiddle := NewMouseAreaBind(256, 0, 256, 256)
	bottomRight := NewMouseAreaBind(512, 0, 256, 256)

	boardbtns := boardButtons{
		topLeft:      topLeft,
		topMiddle:    topMiddle,
		topRight:     topRight,
		middleLeft:   middleLeft,
		middle:       middle,
		middleRight:  middleRight,
		bottomLeft:   bottomLeft,
		bottomMiddle: bottomMiddle,
		bottomRight:  bottomRight,
	}

	for !win.Closed() {
		imd.Clear()
		win.Clear(colornames.Ghostwhite)
		//self explanatory
		if mainmenu == true {
			logicdrawMainMenu(&mainmenu, menutitle, atlas, win, imd, button, &hoverstreamer)
		} else if tie {
			logicdrawTie(&mainmenu, tietext, atlas, win, imd, endbutton, &hoverstreamer, reset)
		} else if won || won2 {
			logicdrawWon(&won, &won2, &mainmenu, wontext, won2text, atlas, win, imd, endbutton, &hoverstreamer, reset)
		} else {
			logicdrawBoard(&tie, &placed, &won, &won2, &mainmenu, &board, win, imd, boardbtns, &starter, &turn, &pos)
		}
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
