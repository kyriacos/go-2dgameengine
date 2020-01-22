package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
)

type TankEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
}

func NewTankEntity(am *core.AssetManager) *TankEntity {
	e := &TankEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(0, 0, 20, 20, 32, 32, 1, e)
	e.SpriteComponent = components.NewSpriteComponent(am, e.TransformComponent, "tank-image", e)

	return e
}

func (e *TankEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
