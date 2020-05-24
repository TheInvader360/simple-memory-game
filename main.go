package main

import (
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

var square *ebiten.Image

func update(screen *ebiten.Image) error {
  ebitenutil.DebugPrint(screen, "Hello Ebiten!")
  return nil
}

func main() {
  if err := ebiten.Run(update, 320, 240, 2, "Hello Ebiten!"); err != nil {
    panic(err)
  }
}
