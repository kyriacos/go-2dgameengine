package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteEntity struct {
	*ecs.Entity
	*components.SpriteComponent
}
type RenderSpritesSystem struct {
	entities []*spriteEntity
}

func (r *RenderSpritesSystem) Add(e *ecs.Entity, sc *components.SpriteComponent) {
	r.entities = append(r.entities, &spriteEntity{Entity: e, SpriteComponent: sc})
}

func (r *RenderSpritesSystem) Update(dt float64) {
	clear() // clear the buffer
	for _, e := range r.entities {
		s := e.SpriteComponent
		if s.IsAnimated {
			s.SourceRectangle.X = s.SourceRectangle.W * int32((int32(sdl.GetTicks())/s.AnimationSpeed)%s.NumFrames)
			s.SourceRectangle.Y = int32(s.AnimationIndex * s.TransformComponent.Height)
		}

		s.DestinationRectangle.X = int32(s.TransformComponent.Position.X)
		s.DestinationRectangle.Y = int32(s.TransformComponent.Position.Y)
		s.DestinationRectangle.W = int32(s.TransformComponent.Width * s.TransformComponent.Scale)
		s.DestinationRectangle.H = int32(s.TransformComponent.Height * s.TransformComponent.Scale)

		core.DrawTexture(s.Texture, s.SourceRectangle, s.DestinationRectangle, s.SpriteFlip)
	}
}
