package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardEntity struct {
	*ecs.Entity
	*components.PlayerControlComponent
}
type PlayerControlSystem struct {
	entities []*keyboardEntity
}

func (s *PlayerControlSystem) Add(e *ecs.Entity, kc *components.PlayerControlComponent) {
	s.entities = append(s.entities, &keyboardEntity{Entity: e, PlayerControlComponent: kc})
}

func (s *PlayerControlSystem) Update(dt float64) {
	clear() // clear the buffer
	for _, e := range s.entities {
		k := e.PlayerControlComponent
		transform := k.TransformComponent
		sprite := k.SpriteComponent

		switch t := global.Event.(type) {
		case *sdl.KeyboardEvent:
			key := t.Keysym.Sym
			if t.Type == sdl.KEYDOWN {
				switch key {
				case sdl.K_UP:
					transform.Velocity.Y = -10
					transform.Velocity.X = 0
					sprite.Play("Up")
				case sdl.K_RIGHT:
					transform.Velocity.Y = 0
					transform.Velocity.X = 10
					sprite.Play("Right")
				case sdl.K_DOWN:
					transform.Velocity.Y = 10
					transform.Velocity.X = 0
					sprite.Play("Down")
				case sdl.K_LEFT:
					transform.Velocity.Y = 0
					transform.Velocity.X = -10
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
	}

}
