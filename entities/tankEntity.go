package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type TankEntity struct {
	*ecs.Entity
	*components.TransformComponent `json:"TransformComponent"`
	*components.SpriteComponent    `json:"SpriteComponent"`
	*components.ColliderComponent  `json:"ColliderComponent"`
}

func NewTankEntity(am *core.AssetManager) *TankEntity {
	e := &TankEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(150, 495, 20, 0, 32, 32, 1, e)
	e.SpriteComponent = components.NewSpriteComponent(am, e.TransformComponent, "tank-image", false, e)
	e.ColliderComponent = components.NewColliderComponent(enums.ColliderTagEnemy, vec.Vector2{X: 150, Y: 495}, 32, 32, e)
	return e
}

func (e *TankEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
