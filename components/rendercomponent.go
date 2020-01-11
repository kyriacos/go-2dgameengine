package components

import (
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

// type Drawable interface {
// 	Texture() *sdl.Texture
// 	Width() float64
// 	Height() float64
// }

type RenderComponent struct {
	*ecs.Component
	Shape sdl.Rect // maybe define an enum instead RECT = 1
}
