package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

type debugEntity struct {
	*ecs.Entity
	*components.ColliderComponent
}

type RenderDebugSystem struct {
	entities []*debugEntity
	AManager *core.AssetManager
}

func (s *RenderDebugSystem) Add(e *ecs.Entity, cc *components.ColliderComponent) {
	s.entities = append(s.entities, &debugEntity{Entity: e, ColliderComponent: cc})
}

func (s *RenderDebugSystem) Update(dt float64) {
	if !global.EnableDebug {
		return
	}
	for _, e := range s.entities {
		s.renderBoundingBox(e)
	}
}

func (s *RenderDebugSystem) renderBoundingBox(e *debugEntity) {
	// //TODO: remove hardcoding
	texture := s.AManager.GetTexture("collision-texture")
	c := e.ColliderComponent
	core.DrawTexture(texture, c.SourceRectangle, c.DestinationRectangle, sdl.FLIP_NONE)
}
