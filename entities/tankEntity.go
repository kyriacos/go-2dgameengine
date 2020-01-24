package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type TankEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.ColliderComponent
}

func NewTankEntity(am *core.AssetManager) *TankEntity {
	e := &TankEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(150, 495, 5, 0, 32, 32, 1, e)
	e.SpriteComponent = components.NewSpriteComponent(am, e.TransformComponent, "tank-image", e)
	e.ColliderComponent = components.NewColliderComponent("enemy", vec.Vector2{X: 150, Y: 495}, 32, 32, e)
	return e
}

func (e *TankEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
