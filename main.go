package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	Window   *sdl.Window
	Renderer *sdl.Renderer

	// WindowWidth  = flag.Int("width", 640, "the window width")
	// WindowHeight = flag.Int("height", 480, "the window height")
	WindowWidth  = 800
	WindowHeight = 600

	running = false
)

func initSDL() (err error) {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing SDL: %s\n", err)
		return err
	}

	Window, err = sdl.CreateWindow(
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

	Renderer, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED) // -1 to use the default graphics driver
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
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.Clear()
}

func update() {

}

func render() {
	Renderer.Present()
}

func destroy() {
	defer sdl.Quit()
	defer Window.Destroy()
	defer Renderer.Destroy()
}

func main() {
	flag.Parse()

	if err := initSDL(); err != nil {
		destroy() // cleanup and exit
		os.Exit(1)
	}

	setup()

	running = true
	for running {
		processInput()
		update()
		render()
	}
}
