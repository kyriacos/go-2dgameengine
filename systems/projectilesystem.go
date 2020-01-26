package systems

import (
	"math"

	"github.com/kyriacos/2dgameengine/core"

	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
)

type projectileEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.SpriteComponent
	*components.ColliderComponent
	*components.ProjectileEmitterComponent
}

type ProjectileSystem struct {
	EM       *core.EntityManager
	entities []*projectileEntity
}

func (s *ProjectileSystem) Add(
	e *ecs.Entity,
	tc *components.TransformComponent,
	sp *components.SpriteComponent,
	cc *components.ColliderComponent,
	pec *components.ProjectileEmitterComponent) {
	s.entities = append(s.entities, &projectileEntity{
		Entity:                     e,
		TransformComponent:         tc,
		SpriteComponent:            sp,
		ColliderComponent:          cc,
		ProjectileEmitterComponent: pec,
	})
}

func (s *ProjectileSystem) Update(dt float64) {
	// loopedEntities := []*projectileEntity{}
	for _, e := range s.entities {
		transform := e.TransformComponent
		pec := e.ProjectileEmitterComponent

		// transform.Velocity = vec.Vector2{X: float32(math.Cos(pec.Angle) * float64(pec.Speed)), Y: float32(math.Sin(pec.Angle) * float64(pec.Speed))}

		if distanceBetweenPoints(transform.Position, pec.Origin) > float64(pec.Range) {
			if pec.Loop {
				transform.Position = pec.Origin
				// loopedEntities = append(loopedEntities, e)
			} else {
				// TODO: destroy the entity. i.e. remove from BOTH the system and entity manager
				// s.Destroy()

			}
		}
	}
	// s.entities = loopedEntities[:]
}

func distanceBetweenPoints(p1, p2 vec.Vector2) float64 {
	return float64(math.Sqrt(float64((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))))
}
