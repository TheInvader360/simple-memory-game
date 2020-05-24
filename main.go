package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenSize, padSize = 480, 220
)

var (
	targetSequence sequence
	pads           []pad
	level          int
	demo           bool
	listening      bool
	tickCounter    int

	blueDark    = &color.NRGBA{0x00, 0x00, 0x33, 0xff}
	blueLight   = &color.NRGBA{0x00, 0x00, 0xff, 0xff}
	greenDark   = &color.NRGBA{0x00, 0x33, 0x00, 0xff}
	greenLight  = &color.NRGBA{0x00, 0xff, 0x00, 0xff}
	redDark     = &color.NRGBA{0x33, 0x00, 0x00, 0xff}
	redLight    = &color.NRGBA{0xff, 0x00, 0x00, 0xff}
	yellowDark  = &color.NRGBA{0x33, 0x33, 0x00, 0xff}
	yellowLight = &color.NRGBA{0xff, 0xff, 0x00, 0xff}
)

type sequence struct {
	values       []int
	currentIndex int
}

type pad struct {
	x, y              float64
	on                bool
	imageOff, imageOn *ebiten.Image
}

func init() {
	rand.Seed(time.Now().UnixNano())

	targetSequence.values = make([]int, 30)
	for index := range targetSequence.values {
		targetSequence.values[index] = rand.Intn(4)
	}
	fmt.Println(targetSequence.values)

	blueDarkImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	blueDarkImage.Fill(blueDark)
	blueLightImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	blueLightImage.Fill(blueLight)
	greenDarkImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	greenDarkImage.Fill(greenDark)
	greenLightImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	greenLightImage.Fill(greenLight)
	redDarkImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	redDarkImage.Fill(redDark)
	redLightImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	redLightImage.Fill(redLight)
	yellowDarkImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	yellowDarkImage.Fill(yellowDark)
	yellowLightImage, _ := ebiten.NewImage(padSize, padSize, ebiten.FilterDefault)
	yellowLightImage.Fill(yellowLight)

	pads = append(pads, pad{0, 0, false, blueDarkImage, blueLightImage})
	pads = append(pads, pad{screenSize - padSize, 0, false, greenDarkImage, greenLightImage})
	pads = append(pads, pad{0, screenSize - padSize, false, redDarkImage, redLightImage})
	pads = append(pads, pad{screenSize - padSize, screenSize - padSize, false, yellowDarkImage, yellowLightImage})

	level = 3

	demo = true
}

func update(screen *ebiten.Image) error {
	tickCounter++

	if demo {
		if targetSequence.currentIndex < level {
			if tickCounter == 1 {
				allPadsOff()
				pads[targetSequence.values[targetSequence.currentIndex]].on = true
			}
			if tickCounter == 41 {
				allPadsOff()
			}
			if tickCounter == 61 {
				tickCounter = 0
				targetSequence.currentIndex++
			}
		} else {
			targetSequence.currentIndex = 0
			demo = false
			listening = true
		}
	}

	if listening {
		allPadsOff()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			posX, posY := ebiten.CursorPosition()
			pos := image.Point{posX, posY}
			for index, pad := range pads {
				if pos.In(pad.imageOff.Bounds().Add(image.Point{int(pad.x), int(pad.y)})) {
					pads[index].on = true
				}
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, pad := range pads {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(pad.x, pad.y)
		if pad.on {
			screen.DrawImage(pad.imageOn, opts)
		} else {
			screen.DrawImage(pad.imageOff, opts)
		}
	}

	return nil
}

func allPadsOff() {
	for index := range pads {
		pads[index].on = false
	}
}

func main() {
	if err := ebiten.Run(update, screenSize, screenSize, 1, "Simple Memory Game"); err != nil {
		panic(err)
	}
}
