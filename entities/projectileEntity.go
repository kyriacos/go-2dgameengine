package entities

import (
	"math"

	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/core/util"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type ProjectileEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.ColliderComponent
	*components.ProjectileEmitterComponent
}

func NewProjectileEntity(am *core.AssetManager) *ProjectileEntity {
	e := &ProjectileEntity{Entity: ecs.NewEntity()}
	// We add 16 since we want it to start after the tank

	e.TransformComponent = components.NewTransformComponent(
		150+16, 495+16, 0, 0, 4, 4, 1, e)
	e.SpriteComponent = components.NewSpriteComponent(
		am, e.TransformComponent, "projectile-image", false, e)
	e.ColliderComponent = components.NewColliderComponent(
		enums.ColliderTagLevelComplete, vec.Vector2{X: 150 + 16, Y: 495 + 16}, 32, 32, e)
	e.ProjectileEmitterComponent = components.NewProjectileEmitterComponent(
		vec.Vector2{X: 150 + 16, Y: 495 + 16}, 50, 200, 270, true, e)

	angleRad := util.DegToRad(float64(270))
	e.TransformComponent.Velocity = vec.Vector2{X: float32(math.Cos(angleRad) * float64(50)), Y: float32(math.Sin(angleRad) * float64(50))}
	return e
}

func (e *ProjectileEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}
