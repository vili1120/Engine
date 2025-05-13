package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vili1120/Engine/scenev"
)

func main() {
  runtime.LockOSThread()
  app := scenev.Init()

  app.InitWindow(800, 600, "i dont know what to put here")
  defer app.CloseWindow()

  for !app.WindowShouldClose() {
    glfw.PollEvents()
  }
}
