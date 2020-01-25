package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/ecs"
)

type moveableEntity struct {
	*ecs.Entity
	*components.TransformComponent
}

type MoveableSystem struct {
	entities []*moveableEntity
}

func (s *MoveableSystem) Add(e *ecs.Entity, tc *components.TransformComponent) {
	s.entities = append(s.entities, &moveableEntity{
		Entity:             e,
		TransformComponent: tc,
	})
}

func (s *MoveableSystem) Update(dt float64) {
	for _, e := range s.entities {
		t := e.TransformComponent
		change := t.Velocity.Mul(float32(dt))
		t.Position = t.Position.Add(change)
	}
}
