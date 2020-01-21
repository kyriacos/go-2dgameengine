package components

import (
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteComponent struct {
	ecs.Component
	Texture              *sdl.Texture
	SourceRectangle      *sdl.Rect
	DestinationRectangle *sdl.Rect
	TransformComponent   *TransformComponent
	SpriteFlip           sdl.RendererFlip
}

// func New(
// 	am *core.AssetManager,
// 	transform *TransformComponent,
// 	textureId string,
// ) {
// 	return &SpriteComponent{
// 		Texture:              texture,
// 		TransformComponent:   transform,
// 		SourceRectangle:      &sdl.Rect{X: 0, Y: 0, W: int32(transform.Width), H: int32(transform.Height)},
// 		DestinationRectangle: &sdl.Rect{},
// 		SpriteFlip:           sdl.FLIP_NONE,
// 	}
// }
// func Init(owner ecs.IEntity) {
// 	// set the new instances or whaterver
// }

func NewSpriteComponent(
	am *core.AssetManager,
	transform *TransformComponent,
	textureId string,
	owner ecs.IEntity,
) *SpriteComponent {

	texture := am.GetTexture(textureId)

	return &SpriteComponent{
		Texture:              texture,
		TransformComponent:   transform,
		SourceRectangle:      &sdl.Rect{X: 0, Y: 0, W: int32(transform.Width), H: int32(transform.Height)},
		DestinationRectangle: &sdl.Rect{},
		SpriteFlip:           sdl.FLIP_NONE,
	}
}

func (c *SpriteComponent) SetTexture(textureId string) {
	// c.texture =
}
