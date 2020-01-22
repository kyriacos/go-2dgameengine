package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
)

type ChopperEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.PlayerControlComponent
}

func NewChopperEntity(am *core.AssetManager) *ChopperEntity {
	e := &ChopperEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(240, 106, 0, 0, 32, 32, 1, e)
	e.SpriteComponent = components.NewAnimatedSpriteComponent(am, e.TransformComponent, "chopper-image", 2, 90, true, false, e)
	e.PlayerControlComponent = components.NewPlayerControlComponent(e.TransformComponent, e.SpriteComponent, e)

	return e
}
