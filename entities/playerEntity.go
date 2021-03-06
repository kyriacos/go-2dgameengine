package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type PlayerEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.CameraComponent
	*components.ColliderComponent
}

func NewPlayerEntity(am *core.AssetManager) *PlayerEntity {
	e := &PlayerEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(240, 106, 0, 0, 32, 32, 1, e)
	e.SpriteComponent = components.NewAnimatedSpriteComponent(am, e.TransformComponent, "player-image", 2, 90, true, false, e)
	e.CameraComponent = components.NewCameraComponent(vec.Vector2{X: 0, Y: 0}, e)
	e.ColliderComponent = components.NewColliderComponent(enums.ColliderTagPlayer, vec.Vector2{X: 240, Y: 106}, 32, 32, e)

	return e
}

func (e *PlayerEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
