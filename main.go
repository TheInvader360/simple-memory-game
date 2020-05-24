package main

import (
	"fmt"
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
	blueDark    = &color.NRGBA{0x00, 0x00, 0x33, 0xff}
	blueLight   = &color.NRGBA{0x00, 0x00, 0xff, 0xff}
	greenDark   = &color.NRGBA{0x00, 0x33, 0x00, 0xff}
	greenLight  = &color.NRGBA{0x00, 0xff, 0x00, 0xff}
	redDark     = &color.NRGBA{0x33, 0x00, 0x00, 0xff}
	redLight    = &color.NRGBA{0xff, 0x00, 0x00, 0xff}
	yellowDark  = &color.NRGBA{0x33, 0x33, 0x00, 0xff}
	yellowLight = &color.NRGBA{0xff, 0xff, 0x00, 0xff}

	targetSequence sequence
	pads           []pad
	level          int
	demo           bool
	tickCounter    int
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

	level = 10

	demo = true
}

func update(screen *ebiten.Image) error {
	tickCounter++

	if demo {
		if targetSequence.currentIndex < level {
			if tickCounter == 1 {
				pads[0].on = false
				pads[1].on = false
				pads[2].on = false
				pads[3].on = false
				pads[targetSequence.values[targetSequence.currentIndex]].on = true
			}
			if tickCounter == 41 {
				pads[0].on = false
				pads[1].on = false
				pads[2].on = false
				pads[3].on = false
			}
			if tickCounter == 61 {
				tickCounter = 0
				targetSequence.currentIndex++
			}
		} else {
			targetSequence.currentIndex = 0
			demo = false
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

func main() {
	if err := ebiten.Run(update, screenSize, screenSize, 1, "Simple Memory Game"); err != nil {
		panic(err)
	}
}
