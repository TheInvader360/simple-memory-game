package main

import (
  "image/color"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

var square *ebiten.Image

func update(screen *ebiten.Image) error {
  screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})
  ebitenutil.DebugPrint(screen, "Hello Ebiten!")
  if square == nil {
    square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
  }
  opts := &ebiten.DrawImageOptions{}
  opts.GeoM.Translate(32, 32)
  square.Fill(color.Black)
  screen.DrawImage(square, opts)
  opts.GeoM.Translate(32, 32)
  opts.GeoM.Scale(2, 2)
  square.Fill(color.White)
  screen.DrawImage(square, opts)
  return nil
}

func main() {
  if err := ebiten.Run(update, 320, 240, 2, "Hello Ebiten!"); err != nil {
    panic(err)
  }
}
