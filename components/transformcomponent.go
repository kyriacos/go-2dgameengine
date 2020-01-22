package components

import (
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type TransformComponent struct {
	*ecs.Component
	Position, Velocity   vec.Vector2
	Width, Height, Scale int32
}

func NewTransformComponent(posX, posY int32, velX, velY float32, w, h, s int32, owner ecs.IEntity) *TransformComponent {
	return &TransformComponent{
		Component: ecs.NewBaseComponent(owner),
		Position:  vec.Vector2{X: float32(posX), Y: float32(posY)},
		Velocity:  vec.Vector2{X: velX, Y: velY},
		Width:     w,
		Height:    h,
		Scale:     s,
	}
}
