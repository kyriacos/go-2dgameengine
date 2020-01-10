package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	Window *sdl.Window

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
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SDL Window: %s\n", err)
		return err
	}

	return nil
}

func destroy() {
	defer sdl.Quit()
	defer Window.Destroy()
}

func main() {
	flag.Parse()

	if err := initSDL(); err != nil {
		destroy() // cleanup and exit
		os.Exit(1)
	}

	running = true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
