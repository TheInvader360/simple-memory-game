package main

import (
  "fmt"
  "image/color"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

var square *ebiten.Image

func update(screen *ebiten.Image) error {
  screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})
  x, y := ebiten.CursorPosition()
  ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))
  if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
    ebitenutil.DebugPrint(screen, "\n\nTry pressing LEFT MOUSE BUTTON or RIGHT MOUSE BUTTON")
  }
  if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
      ebitenutil.DebugPrint(screen, "\n\n\nPressing LEFT MOUSE BUTTON")
  }
  if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
    ebitenutil.DebugPrint(screen, "\n\n\n\nPressing RIGHT MOUSE BUTTON")
  }
  if !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
    ebitenutil.DebugPrint(screen, "\n\n\n\n\n\nTry pressing UP/DOWN/LEFT/RIGHT cursor keys")
  }
  if ebiten.IsKeyPressed(ebiten.KeyUp) {
    ebitenutil.DebugPrint(screen, "\n\n\n\n\n\n\nPressing UP")
  }
  if ebiten.IsKeyPressed(ebiten.KeyDown) {
    ebitenutil.DebugPrint(screen, "\n\n\n\n\n\n\n\nPressing DOWN")
  }
  if ebiten.IsKeyPressed(ebiten.KeyLeft) {
    ebitenutil.DebugPrint(screen, "\n\n\n\n\n\n\n\n\nPressing LEFT")
  }
  if ebiten.IsKeyPressed(ebiten.KeyRight) {
    ebitenutil.DebugPrint(screen, "\n\n\n\n\n\n\n\n\n\nPressing RIGHT")
  }
  if square == nil {
    square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
  }
  opts := &ebiten.DrawImageOptions{}
  opts.GeoM.Translate(128, 128)
  square.Fill(color.Black)
  screen.DrawImage(square, opts)
  opts.GeoM.Translate(-32, -64)
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
