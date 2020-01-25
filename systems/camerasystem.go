package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core/util"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/kyriacos/2dgameengine/vec"
)

type cameraEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.CameraComponent
}

type CameraSystem struct {
	entities []*cameraEntity
}

func (s *CameraSystem) Add(e *ecs.Entity,
	transform *components.TransformComponent,
	cc *components.CameraComponent) {
	s.entities = append(s.entities, &cameraEntity{
		Entity:             e,
		TransformComponent: transform,
		CameraComponent:    cc,
	})
}

func (s *CameraSystem) Update(dt float64) {
	for _, e := range s.entities {
		transform := e.TransformComponent
		camera := e.CameraComponent

		x := int(transform.Position.X) - global.WindowWidth/2
		y := int(transform.Position.Y) - global.WindowHeight/2

		camera.Position = vec.Vector2{
			X: float32(util.Clamp(x, 0, int(camera.Width))),
			Y: float32(util.Clamp(y, 0, int(camera.Height))),
		}
	}
}
