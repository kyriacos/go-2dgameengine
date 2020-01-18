package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/entities"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/kyriacos/2dgameengine/systems"
)

var (
	World *ecs.World

	// WindowWidth  = flag.Int("width", 640, "the window width")
	// WindowHeight = flag.Int("height", 480, "the window height")
	WindowWidth  = 800
	WindowHeight = 600

	running = false

	showFPS = flag.Bool("showFPS", false, "Show current FPS and on exit display the average FPS.")
)

func initSDL() (err error) {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing SDL: %s\n", err)
		return err
	}

	global.Window, err = sdl.CreateWindow(
		"",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(WindowWidth),
		int32(WindowHeight),
		sdl.WINDOW_BORDERLESS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SDL Window: %s\n", err)
		return err
	}

	global.Renderer, err = sdl.CreateRenderer(global.Window, -1, sdl.RENDERER_ACCELERATED) // -1 to use the default graphics driver
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SDL renderer: %s\n", err)
		return err
	}

	// if err = Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to set blend mode: %s", err)
	// 	return
	// }

	return nil
}

func processInput() {
	if event := sdl.PollEvent(); event != nil {
		switch t := event.(type) {
		case *sdl.QuitEvent: // sdl.QUIT
			running = false
		case *sdl.KeyboardEvent:
			key := t.Keysym.Sym
			if t.Type == sdl.KEYDOWN {
				switch key {
				case sdl.K_ESCAPE:
					running = false
				}
			}
		}
	}
}

func setup() {
	global.Renderer.SetDrawColor(0, 0, 0, 255)
	global.Renderer.Clear()

	// Game - Load Level

	// Create Entity Manager
	em := &core.EntityManager{
		Entities: make(map[uint64]ecs.IEntity),
	}

	// Create asset manager
	am := &core.AssetManager{
		EntityManager: em,
		Textures:      make(map[string]*sdl.Texture),
	}
	// level 0
	textureFilePath := "./assets/images/tank-big-right.png"
	am.AddTexture("tank-image", textureFilePath)

	// create entities and components
	e := entities.NewTankEntity(am)
	// Add to entitymanager
	em.AddEntity(e)

	// create world
	World = &ecs.World{}
	// add systems
	// renderSystem := &systems.RenderSystem{}
	// renderSystem.Add(pe.Entity, pe.RenderComponent, pe.TransformComponent)
	renderSpritesSystem := &systems.RenderSpritesSystem{}
	renderSpritesSystem.Add(e.Entity, e.SpriteComponent)

	moveableSystem := &systems.MoveableSystem{}
	moveableSystem.Add(e.Entity, e.TransformComponent)

	World.AddSystem(moveableSystem)
	World.AddSystem(renderSpritesSystem)
}

func update(deltaTime float64) {
	// Game world update
	World.Update(deltaTime)
}

func render() {
	global.Renderer.Present()
}

func destroy() {
	defer sdl.Quit()
	defer global.Window.Destroy()
	defer global.Renderer.Destroy()
}

func main() {
	flag.Parse()

	if err := initSDL(); err != nil {
		destroy() // cleanup and exit
		os.Exit(1)
	}

	var (
		counter                      = 0
		sumFPS, deltaTime            float64
		ticksCurrent, ticksLastFrame uint32
	)

	setup()

	running = true
	for running {
		ticksCurrent = sdl.GetTicks()

		deltaTime = float64(ticksCurrent-ticksLastFrame) / 1000.0

		deltaTime = math.Min(deltaTime, 0.05) // clamp deltatime to max 0.05

		processInput()
		update(deltaTime)
		render()

		ticksLastFrame = sdl.GetTicks()

		sdl.Delay(uint32(math.Floor(FrameTargetTime - deltaTime))) // pause until we reach the target frames

		if *showFPS {
			counter++
			currentFPS := 1.0 / deltaTime
			sumFPS += currentFPS

			fmt.Printf("FPS: %f\n", currentFPS)
		}
	}

	destroy()

	if *showFPS {
		fmt.Printf("Average FPS: %f\n", sumFPS/float64(counter))
	}

	os.Exit(0)

}
