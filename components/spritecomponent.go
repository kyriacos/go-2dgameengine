package components

import (
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type Animation struct {
	index          int32
	numFrames      int32
	animationSpeed int32
}

type SpriteComponent struct {
	*ecs.Component
	Texture              *sdl.Texture
	SourceRectangle      *sdl.Rect
	DestinationRectangle *sdl.Rect
	TransformComponent   *TransformComponent
	SpriteFlip           sdl.RendererFlip

	IsAnimated     bool
	IsFixed        bool // doesn't move
	NumFrames      int32
	AnimationSpeed int32

	Animations           map[string]Animation
	AnimationIndex       int32
	CurrentAnimationName string
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
	isFixed bool,
	owner ecs.IEntity,
) *SpriteComponent {

	texture := am.GetTexture(textureId)

	return &SpriteComponent{
		Component:            ecs.NewBaseComponent(owner),
		Texture:              texture,
		TransformComponent:   transform,
		SourceRectangle:      &sdl.Rect{X: 0, Y: 0, W: int32(transform.Width), H: int32(transform.Height)},
		DestinationRectangle: &sdl.Rect{},
		SpriteFlip:           sdl.FLIP_NONE,
		IsFixed:              isFixed,
		IsAnimated:           false,
	}
}

func NewAnimatedSpriteComponent(
	am *core.AssetManager,
	transform *TransformComponent,
	textureID string,
	numFrames int32,
	animationSpeed int32,
	hasDirections bool,
	isFixed bool,
	owner ecs.IEntity,
) *SpriteComponent {

	texture := am.GetTexture(textureID)
	c := &SpriteComponent{
		Component:            ecs.NewBaseComponent(owner),
		Texture:              texture,
		TransformComponent:   transform,
		SourceRectangle:      &sdl.Rect{X: 0, Y: 0, W: int32(transform.Width), H: int32(transform.Height)},
		DestinationRectangle: &sdl.Rect{},
		SpriteFlip:           sdl.FLIP_NONE,
		NumFrames:            numFrames,
		AnimationSpeed:       animationSpeed,
		IsAnimated:           true,
		IsFixed:              isFixed,
	}

	if hasDirections {
		down := Animation{0, numFrames, animationSpeed}
		right := Animation{1, numFrames, animationSpeed}
		left := Animation{2, numFrames, animationSpeed}
		up := Animation{3, numFrames, animationSpeed}
		c.Animations = make(map[string]Animation, 4)
		c.Animations["Down"] = down
		c.Animations["Right"] = right
		c.Animations["Left"] = left
		c.Animations["Up"] = up
		c.AnimationIndex = 0
		c.CurrentAnimationName = "Down"
	} else {
		animation := Animation{0, numFrames, animationSpeed}
		c.Animations = make(map[string]Animation, 1)
		c.Animations["SingleAnimation"] = animation
		c.AnimationIndex = 0
		c.CurrentAnimationName = "SingleAnimation"
	}

	c.Play(c.CurrentAnimationName)

	return c
}

func (c *SpriteComponent) Play(animationName string) {
	c.NumFrames = c.Animations[animationName].numFrames
	c.AnimationIndex = c.Animations[animationName].index
	c.AnimationSpeed = c.Animations[animationName].animationSpeed
	c.CurrentAnimationName = animationName
}

func (c *SpriteComponent) SetTexture(textureId string) {
	// c.texture =
}
