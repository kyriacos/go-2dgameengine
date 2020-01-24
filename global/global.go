package global

import "github.com/veandco/go-sdl2/sdl"

var (
	Window   *sdl.Window
	Renderer *sdl.Renderer

	Event sdl.Event

	WindowWidth, WindowHeight int

	Running = false
)
