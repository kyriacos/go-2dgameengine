package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
)

type RadarEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
}

func NewRadarEntity(am *core.AssetManager) *RadarEntity {
	e := &RadarEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(720, 15, 0, 0, 64, 64, 1, e)
	e.SpriteComponent = components.NewAnimatedSpriteComponent(am, e.TransformComponent, "radar-image", 8, 150, false, true, e)

	return e
}

func (e *RadarEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
