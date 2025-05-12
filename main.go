package main

import (
  "runtime"
	"github.com/vili1120/Engine/scenev"
)

func main() {
  runtime.LockOSThread()
  scenev.InitWindow(800, 600, "title")
  for !scenev.WindowShouldClose() {
    continue
  }
}
