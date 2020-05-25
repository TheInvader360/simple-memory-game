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
	screenSize, padSize, minLevel, maxLevel, newLevelDelayTicks, lightTicks, darkTicks = 480, 220, 1, 10, 60, 40, 20
)

var (
	sequence                         []int
	pads                             []pad
	lastPressedPad                   *pad
	demoMode, playMode               bool
	level, currentIndex, tickCounter int

	blueDark    = &color.NRGBA{0x00, 0x00, 0x33, 0xff}
	blueLight   = &color.NRGBA{0x00, 0x00, 0xff, 0xff}
	greenDark   = &color.NRGBA{0x00, 0x33, 0x00, 0xff}
	greenLight  = &color.NRGBA{0x00, 0xff, 0x00, 0xff}
	redDark     = &color.NRGBA{0x33, 0x00, 0x00, 0xff}
	redLight    = &color.NRGBA{0xff, 0x00, 0x00, 0xff}
	yellowDark  = &color.NRGBA{0x33, 0x33, 0x00, 0xff}
	yellowLight = &color.NRGBA{0xff, 0xff, 0x00, 0xff}
)

type pad struct {
	x, y              float64
	on                bool
	imageOff, imageOn *ebiten.Image
}

func init() {
	rand.Seed(time.Now().UnixNano())

	sequence = make([]int, maxLevel)
	for index := range sequence {
		sequence[index] = rand.Intn(4)
	}

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

	level = minLevel - 1
	nextLevel()
}

func update(screen *ebiten.Image) error {
	if demoMode {
		tickCounter++
		if currentIndex < level {
			if tickCounter == 1 {
				allPadsOff()
				pads[sequence[currentIndex]].on = true
			}
			if tickCounter == 1+lightTicks {
				allPadsOff()
			}
			if tickCounter == 1+lightTicks+darkTicks {
				tickCounter = 0
				currentIndex++
			}
		} else {
			currentIndex = 0
			demoMode = false
			playMode = true
		}
	}

	if playMode {
		allPadsOff()
		triggerPad := -1
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			posX, posY := ebiten.CursorPosition()
			pos := image.Point{posX, posY}
			pad := getPadAtPos(&pos)
			if pad != nil {
				pad.on = true
				lastPressedPad = pad
			} else {
				triggerPad = releaseLastPressedPad()
			}
		} else {
			triggerPad = releaseLastPressedPad()
		}

		if triggerPad >= 0 {
			if sequence[currentIndex] == triggerPad {
				if currentIndex+1 < level {
					currentIndex++
				} else {
					nextLevel()
				}
			} else {
				gameOver(fmt.Sprintf("LOST AT LEVEL %d", level))
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		gameOver("QUIT")
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

func getPadAtPos(pos *image.Point) *pad {
	for index, pad := range pads {
		if pos.In(pad.imageOff.Bounds().Add(image.Point{int(pad.x), int(pad.y)})) {
			return &pads[index]
		}
	}
	return nil
}

func releaseLastPressedPad() int {
	returnVal := -1
	if lastPressedPad != nil {
		for index, pad := range pads {
			if pad == *lastPressedPad {
				returnVal = index
			}
		}
		lastPressedPad = nil
	}
	return returnVal
}

func nextLevel() {
	if level == maxLevel {
		gameOver("COMPLETED")
	} else {
		level++
		currentIndex = 0
		tickCounter = -newLevelDelayTicks
		demoMode = true
		playMode = false
		fmt.Println(fmt.Sprintf("LEVEL %v", level))
	}
}

func gameOver(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func main() {
	if err := ebiten.Run(update, screenSize, screenSize, 1, "Simple Memory Game"); err != nil {
		panic(err)
	}
}
