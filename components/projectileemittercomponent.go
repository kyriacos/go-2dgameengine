package components

import (
	"github.com/kyriacos/2dgameengine/core/util"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type ProjectileEmitterComponent struct {
	*ecs.Component
	Origin       vec.Vector2
	Speed, Range int32
	Angle        float64
	Loop         bool
}

func NewProjectileEmitterComponent(position vec.Vector2, speed, projRange, angleDegrees int32, loop bool, owner ecs.IEntity) *ProjectileEmitterComponent {
	angleRad := util.DegToRad(float64(angleDegrees))

	return &ProjectileEmitterComponent{
		Component: ecs.NewBaseComponent(owner),
		Origin:    position,
		Speed:     speed,
		Range:     projRange,
		Angle:     angleRad,
		Loop:      loop,
	}
	// transform.velocty = vec.Vector2{X: math.Cos(angleRad) * speed, Y: math.Sin(angleRad) * speed}
}
