package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type TextLabelEntity struct {
	*ecs.Entity
	*components.TextLabelComponent
}

func NewTextLabelEntity(
	am *core.AssetManager,
	x, y int32,
	text, fontFamily string,
	color *sdl.Color,
) *TextLabelEntity {
	e := &TextLabelEntity{Entity: ecs.NewEntity()}
	e.TextLabelComponent = components.NewTextLabelComponent(am, x, y, text, fontFamily, color, e)

	return e
}

func (e *TextLabelEntity) RenderType() ecs.Renderable {
	return e.TextLabelComponent
}
