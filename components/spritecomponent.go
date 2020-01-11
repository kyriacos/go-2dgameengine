package components

import (
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteComponent struct {
	ecs.Component
	texture              *sdl.Texture
	sourceRectangle      *sdl.Rect
	destinationRectangle *sdl.Rect
	spriteFlip           sdl.RendererFlip
}

func NewSpriteComponent(owner *ecs.IEntity) *SpriteComponent {
	transformComponent := owner.
	return &SpriteComponent{
		// texture:              t,
		sourceRectangle:      s,
		destinationRectangle: d,
		spriteFlip:           sdl.FLIP_NONE,
	}
}

func (c *SpriteComponent) SetTexture(textureId string) {
	// c.texture =
}
