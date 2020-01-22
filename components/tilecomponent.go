package components

import (
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/vec"
	"github.com/veandco/go-sdl2/sdl"
)

type TileComponent struct {
	ecs.Component
	Position             vec.Vector2
	Texture              *sdl.Texture
	SourceRectangle      *sdl.Rect // NULL/nil for the entire texture
	DestinationRectangle *sdl.Rect // NULL/nil for the entire texture
	Center               *sdl.Point
	TileFlip             sdl.RendererFlip
	Angle                float32
}

// sdlDestroyTexture - Cleanup somewhere
func NewTileComponent(
	am *core.AssetManager,
	textureID string,
	sourceRectX, sourceRectY,
	x, y,
	tileSize, tileScale int32,
	owner ecs.IEntity,
) *TileComponent {

	texture := am.GetTexture(textureID)

	c := &TileComponent{
		Texture:              texture,
		Position:             vec.Vector2{X: float32(x), Y: float32(y)},
		SourceRectangle:      &sdl.Rect{X: sourceRectX, Y: sourceRectY, W: tileSize, H: tileSize},
		DestinationRectangle: &sdl.Rect{X: x, Y: y, W: tileSize * tileScale, H: tileSize * tileScale},
		TileFlip:             sdl.FLIP_NONE,
	}

	return c
}
