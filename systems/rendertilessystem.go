package systems

import (
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/entities"
)

// type tileEntity struct {
// 	*ecs.Entity
// 	*components.TileComponent
// }
type RenderTilesSystem struct {
	entities []*entities.TileEntity
}

// func (r *RenderTilesSystem) Add(e *ecs.Entity, sc *components.TileComponent) {
// 	r.entities = append(r.entities, &tileEntity{Entity: e, TileComponent: sc})
// }

func (r *RenderTilesSystem) SetEntities(entities []*entities.TileEntity) {
	// r.entities = append(r.entities, &tileEntity{Entity: e, TileComponent: sc})
	r.entities = entities
}

func (r *RenderTilesSystem) Update(dt float64) {
	// clear()
	for _, e := range r.entities {
		s := e.TileComponent

		core.DrawTexture(s.Texture, s.SourceRectangle, s.DestinationRectangle, s.TileFlip)
	}
}
