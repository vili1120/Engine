package scenev

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vulkan-go/vulkan"
)

func InitWindow(w, h int, title string) {
  InitMSG("Initializing SDL2...")
  if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
    ErrorMSG("Failed to initialize SDL2: "+err.Error())
    os.Exit(0)
  }
  defer sdl.Quit()
  SuccessMSG("Initalized SDL2 successfully")
  
  InitMSG("Creating the window")
  window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, int32(w), int32(h), sdl.WINDOW_VULKAN|sdl.WINDOW_SHOWN)

  if err != nil {
    ErrorMSG("Failed to create the window: "+err.Error())
  }
  defer window.Destroy()

  InitMSG("Initializing Vulkan instance...")
	if err := vulkan.SetDefaultGetInstanceProcAddr(); err != nil {
		panic(err)
	}
	if err := vulkan.Init(); err != nil {
		panic(err)
	}
  
  InitMSG("Getting Vulkan instance extensions")
	exts, err := window.VulkanGetInstanceExtensions()
	if err != nil {
		panic(err)
	}

	extNames := make([]*byte, len(exts))
	for i, ext := range exts {
		extNames[i] = vulkan.Str(ext)
	}

	appInfo := vulkan.ApplicationInfo{
		SType:              vulkan.StructureTypeApplicationInfo,
		PApplicationName:   vulkan.Str("Go Vulkan App"),
		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
		PEngineName:        vulkan.Str("No Engine"),
		EngineVersion:      vulkan.MakeVersion(1, 0, 0),
		ApiVersion:         vulkan.API_VERSION_1_0,
	}

	createInfo := vulkan.InstanceCreateInfo{
		SType:                   vulkan.StructureTypeInstanceCreateInfo,
		PApplicationInfo:        &appInfo,
		EnabledExtensionCount:   uint32(len(extNames)),
		PpEnabledExtensionNames: extNames,
	}
  
  InitMSG("Creating Vulkan instance")
	var instance vulkan.Instance
	if vulkan.CreateInstance(&createInfo, nil, &instance) != vulkan.Success {
		ErrorMSG("Failed to create Vulkan instance")
	}
  SuccessMSG("Successfully created Vulkan instance")
	defer vulkan.DestroyInstance(instance, nil)
}

func WindowShouldClose() bool {
  if event := sdl.PollEvent(); event != nil {
    switch event.(type) {
    case *sdl.QuitEvent:
      return true
    }
  }
  return false
}
