package components

import (
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
	"github.com/veandco/go-sdl2/sdl"
)

type ColliderComponent struct {
	*ecs.Component
	Tag                  enums.ColliderTag
	Collider             *sdl.Rect
	SourceRectangle      *sdl.Rect
	DestinationRectangle *sdl.Rect
}

func NewColliderComponent(
	tag enums.ColliderTag,
	position vec.Vector2,
	width, height int32,
	owner ecs.IEntity) *ColliderComponent {

	collider := &sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: width, H: height}

	// dst := collider
	//https://stackoverflow.com/questions/51635766/how-do-i-copy-a-struct-in-golang
	return &ColliderComponent{
		Component:            ecs.NewBaseComponent(owner),
		Tag:                  tag,
		Collider:             collider,
		SourceRectangle:      &sdl.Rect{X: 0, Y: 0, W: width, H: height},
		DestinationRectangle: &sdl.Rect{X: collider.X, Y: collider.Y, W: collider.W, H: collider.H},
	}
}
