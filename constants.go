package main

import "github.com/veandco/go-sdl2/sdl"

const (
	FPS             = 60
	FrameTargetTime = 1000 / FPS
)

var (
	// Colors
	ColorWhite = &sdl.Color{255, 255, 255, 255}
	ColorGreen = &sdl.Color{0, 255, 0, 255}
)
