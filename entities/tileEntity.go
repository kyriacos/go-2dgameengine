package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
)

type TileEntity struct {
	*ecs.Entity
	*components.TileComponent
}

func NewTileEntity(
	am *core.AssetManager,
	textureID string,
	sourceX, sourceY,
	x, y,
	tileSize, scale int32) *TileEntity {

	e := &TileEntity{Entity: ecs.NewEntity()}

	e.TileComponent = components.NewTileComponent(
		am,
		textureID,
		sourceX,
		sourceY,
		x,
		y,
		tileSize,
		scale,
		e,
	)

	return e
}

func (e *TileEntity) RenderType() ecs.Renderable {
	return e.TileComponent
}
