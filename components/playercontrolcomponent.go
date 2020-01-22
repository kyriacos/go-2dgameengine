package components

import (
	"github.com/kyriacos/2dgameengine/ecs"
)

type PlayerControlComponent struct {
	*ecs.Component

	TransformComponent *TransformComponent
	SpriteComponent    *SpriteComponent
}

func NewPlayerControlComponent(
	transform *TransformComponent,
	sprite *SpriteComponent,
	owner ecs.IEntity) *PlayerControlComponent {

	return &PlayerControlComponent{
		Component:          ecs.NewBaseComponent(owner),
		TransformComponent: transform,
		SpriteComponent:    sprite,
	}
}
