package components

import (
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/kyriacos/2dgameengine/vec"
)

type CameraComponent struct {
	*ecs.Component
	Position vec.Vector2
	// Ca:      &sdl.Rect{X: 0, Y: 0, W: int32(transform.Width), H: int32(transform.Height)},
	Scale         int
	Width, Height int32
}

func NewCameraComponent(
	position vec.Vector2,
	owner ecs.IEntity) *CameraComponent {

	return &CameraComponent{
		Component: ecs.NewBaseComponent(owner),
		Position:  position,
		Width:     int32(global.WindowWidth),
		Height:    int32(global.WindowHeight),
		Scale:     2,
	}
}
