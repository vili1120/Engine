package scenev

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	dt "github.com/vili1120/Engine/scenev/devtools"
	"github.com/vulkan-go/vulkan"
)

type App struct {
  window *glfw.Window
}

func Init() *App {
  dt.InitMSG("Initializing GLFW...")
  if err := glfw.Init(); err != nil {
    dt.ErrorMSG("Failed to initialize GLFW: ", err)
  }
  dt.SuccessMSG("Finished intializing GLFW")

  dt.InitMSG("Initializing Vulkan...")
  if err := vulkan.SetDefaultGetInstanceProcAddr(); err != nil {
    dt.ErrorMSG("Failed to run SetDefaultGetInstanceProcAddr: ", err)
  }

  if err := vulkan.Init(); err != nil {
    dt.ErrorMSG("Failed to initialize Vulkan: ", err)
  }
  dt.SuccessMSG("Finished initializing Vulkan")

  return &App{}
}

func (app *App) InitWindow(w, h int32, title string) {
  dt.InitMSG("Creating the window:")
  dt.InitMSG(fmt.Sprintf("  Width: %v", w))
  dt.InitMSG(fmt.Sprintf("  Height: %v", h))
  dt.InitMSG(fmt.Sprintf("  Title: %v", title))

  glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
  
  window, err := glfw.CreateWindow(int(w), int(h), title, nil, nil)
  if err != nil {
    dt.ErrorMSG("Failed to create the window: ", err)
  }
  
  window.Show()

  app.window = window
  dt.SuccessMSG("Finished creating the window")
}

func (app *App) CloseWindow() {
  glfw.Terminate()
  app.window.Destroy()
}

func (app *App) WindowShouldClose() bool {
  return app.window.ShouldClose()
}
