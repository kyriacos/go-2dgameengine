package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core/util"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardEntity struct {
	*ecs.Entity
	*components.SpriteComponent
	*components.TransformComponent
}
type PlayerControlSystem struct {
	entities []*keyboardEntity
}

func (s *PlayerControlSystem) Add(e *ecs.Entity, tc *components.TransformComponent, sp *components.SpriteComponent) {
	s.entities = append(s.entities, &keyboardEntity{
		Entity:             e,
		TransformComponent: tc,
		SpriteComponent:    sp})
}

func (s *PlayerControlSystem) Update(dt float64) {
	for _, e := range s.entities {
		transform := e.TransformComponent
		sprite := e.SpriteComponent

		switch t := global.Event.(type) {
		case *sdl.KeyboardEvent:
			key := t.Keysym.Sym
			if t.Type == sdl.KEYDOWN {
				switch key {
				case sdl.K_UP:
					transform.Velocity.Y = -100
					transform.Velocity.X = 0
					sprite.Play("Up")
				case sdl.K_RIGHT:
					transform.Velocity.Y = 0
					transform.Velocity.X = 100
					sprite.Play("Right")
				case sdl.K_DOWN:
					transform.Velocity.Y = 100
					transform.Velocity.X = 0
					sprite.Play("Down")
				case sdl.K_LEFT:
					transform.Velocity.Y = 0
					transform.Velocity.X = -100
					sprite.Play("Left")
				case sdl.K_SPACE:
					// transform.Velocity.Y = 0
					// transform.Velocity.X = -1
					// sprite.Play("Left")
				}
			}
			if t.Type == sdl.KEYUP {
				switch key {
				case sdl.K_UP:
					transform.Velocity.Y = 0
					// sprite.Play("Up")
				case sdl.K_RIGHT:
					transform.Velocity.X = 0
					// sprite.Play("Right")
				case sdl.K_DOWN:
					transform.Velocity.Y = 0
					// sprite.Play("Down")
				case sdl.K_LEFT:
					transform.Velocity.X = 0
					// sprite.Play("Left")
				case sdl.K_SPACE:
				}
			}

		}

		// Update position
		change := transform.Velocity.Mul(float32(dt))
		position := transform.Position.Add(change)

		// Clamp so it doesn't go off the screen
		position.X = float32(util.Clamp(int(position.X), 0, global.WindowWidth))
		position.Y = float32(util.Clamp(int(position.Y), 0, global.WindowHeight))
		transform.Position = position
	}
}
