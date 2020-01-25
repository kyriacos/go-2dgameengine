package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/entities"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/kyriacos/2dgameengine/systems"
)

var (
	World *ecs.World

	showFPS = false
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
		int32(global.WindowWidth),
		int32(global.WindowHeight),
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

	// if err = global.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to set blend mode: %s", err)
	// 	return
	// }

	return nil
}

func processInput() {
	if global.Event = sdl.PollEvent(); global.Event != nil {
		switch t := global.Event.(type) {
		case *sdl.QuitEvent: // sdl.QUIT
			global.Running = false
		case *sdl.KeyboardEvent:
			key := t.Keysym.Sym
			if t.Type == sdl.KEYDOWN {
				switch key {
				case sdl.K_ESCAPE:
					global.Running = false
				case sdl.K_d:
					global.EnableDebug = !global.EnableDebug
				}
			}
		}
	}
}

func setup() {
	// global.Renderer.SetDrawColor(0, 0, 0, 255)
	// global.Renderer.Clear()

	// Create Entity Manager
	em := &core.EntityManager{
		Entities:      make(map[uint64]ecs.IEntity),
		LayerEntities: make(map[enums.LayerType][]ecs.IEntity),
	}

	// Create asset manager
	am := &core.AssetManager{
		EntityManager: em,
		Textures:      make(map[string]*sdl.Texture),
	}

	// GAME - LOAD LEVEL
	// level 0
	am.AddTexture("tank-image", "./assets/images/tank-big-right.png")           // tank
	am.AddTexture("player-image", "./assets/images/chopper-spritesheet.png")    // player
	am.AddTexture("radar-image", "./assets/images/radar-spritesheet.png")       // radar
	am.AddTexture("helipad-image", "./assets/images/helipad.png")               // radar
	am.AddTexture("jungle-tile-texture", "./assets/tilemaps/jungle.png")        // tiles
	am.AddTexture("collision-texture", "./assets/images/collision-texture.png") // collision bounding box texture

	// game map
	gameMap := NewGameMap(em, am, "jungle-tile-texture", 2, 32)
	gameMap.LoadMap("./assets/tilemaps/jungle.map", 25, 20)

	// tank entity
	tank := entities.NewTankEntity(am)
	em.AddEntity(tank, enums.EnemyLayer)
	// player entity
	player := entities.NewPlayerEntity(am)
	em.AddEntity(player, enums.PlayerLayer)
	// radar entity
	radar := entities.NewRadarEntity(am)
	em.AddEntity(radar, enums.UILayer)
	// level complete entity
	helipad := entities.NewHelipadEntity(am)
	em.AddEntity(helipad, enums.ObstacleLayer)

	// CREATE WORLD
	World = &ecs.World{}

	// ADD SYSTEMS
	pcSystem := &systems.PlayerControlSystem{}
	pcSystem.Add(player.Entity, player.PlayerControlComponent)

	moveableSystem := &systems.MoveableSystem{}
	moveableSystem.Add(tank.Entity, tank.TransformComponent)
	moveableSystem.Add(player.Entity, player.TransformComponent)

	cameraSystem := &systems.CameraSystem{}
	cameraSystem.Add(player.Entity, player.TransformComponent, player.CameraComponent)

	collisionSystem := &systems.CollisionSystem{Camera: player.CameraComponent, AManager: am}
	collisionSystem.Add(player.Entity, player.TransformComponent, player.ColliderComponent)
	collisionSystem.Add(tank.Entity, tank.TransformComponent, tank.ColliderComponent)
	collisionSystem.Add(helipad.Entity, helipad.TransformComponent, helipad.ColliderComponent)

	renderBaseSystem := &systems.RenderBase{}
	renderLayersSystem := &systems.RenderLayersSystem{EM: em, Camera: player.CameraComponent}

	renderDebugSystem := &systems.RenderDebugSystem{AManager: am}
	renderDebugSystem.Add(player.Entity, player.ColliderComponent)
	renderDebugSystem.Add(tank.Entity, tank.ColliderComponent)

	World.AddSystem(moveableSystem)
	World.AddSystem(pcSystem)
	World.AddSystem(cameraSystem)
	World.AddSystem(collisionSystem)
	World.AddSystem(renderBaseSystem)
	World.AddSystem(renderLayersSystem)
	World.AddSystem(renderDebugSystem)
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
	flag.IntVar(&global.WindowWidth, "width", 800, "the window width")
	flag.IntVar(&global.WindowHeight, "height", 600, "the window height")
	flag.BoolVar(&showFPS, "showFPS", false, "Show current FPS and on exit display the average FPS.")

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

	global.Running = true
	for global.Running {
		ticksCurrent = sdl.GetTicks()

		deltaTime = float64(ticksCurrent-ticksLastFrame) / 1000.0

		deltaTime = math.Min(deltaTime, 0.05) // clamp deltatime to max 0.05

		processInput()
		update(deltaTime)
		render()

		ticksLastFrame = sdl.GetTicks()

		sdl.Delay(uint32(math.Floor(FrameTargetTime - deltaTime))) // pause until we reach the target frames

		if showFPS {
			counter++
			currentFPS := 1.0 / deltaTime
			sumFPS += currentFPS

			fmt.Printf("FPS: %f\n", currentFPS)
		}
	}

	destroy()

	if showFPS {
		fmt.Printf("Average FPS: %f\n", sumFPS/float64(counter))
	}

	os.Exit(0)

}
