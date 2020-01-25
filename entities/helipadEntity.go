package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type HelipadEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.ColliderComponent
}

func NewHelipadEntity(am *core.AssetManager) *HelipadEntity {
	e := &HelipadEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(470, 495, 0, 0, 32, 32, 1, e)
	e.SpriteComponent = components.NewSpriteComponent(am, e.TransformComponent, "helipad-image", false, e)
	e.ColliderComponent = components.NewColliderComponent(enums.ColliderTagLevelComplete, vec.Vector2{X: 470, Y: 420}, 32, 32, e)

	return e
}

func (e *HelipadEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
