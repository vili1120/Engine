package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	sv "github.com/vili1120/Engine/SceneV"
)

type Game struct {}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw() {
  rl.ClearBackground(rl.RayWhite)
}

func (g *Game) Layout() (int, int) {
  return 800, 600
}

func main() {
  if err := sv.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }
}
