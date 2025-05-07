package scenev

import rl "github.com/gen2brain/raylib-go/raylib"

type Game interface {
  Draw()
  Update() error
  Layout() (int, int)
}

func RunGame(g Game) error {
  Width, Height := g.Layout()
  rl.InitWindow(int32(Width), int32(Height), "test")
  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {
    g.Update()
    rl.BeginDrawing()
    g.Draw()
    rl.EndDrawing()
  }
  rl.CloseWindow()
  return nil
}
