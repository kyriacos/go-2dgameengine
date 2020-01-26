package components

import (
	"log"

	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

type TextLabelComponent struct {
	*ecs.Component
	Texture    *sdl.Texture
	Position   *sdl.Rect
	Color      *sdl.Color
	Text       string
	FontFamily string
}

func NewTextLabelComponent(
	am *core.AssetManager,
	x, y int32,
	text, fontFamily string,
	color *sdl.Color,
	owner ecs.IEntity) *TextLabelComponent {

	font := am.GetFont(fontFamily)
	surface, err := font.RenderUTF8Blended(text, *color)
	if err != nil {
		log.Fatalf("Could not create UTF8 surface for text label component, fontFamily: %s, text: %s. Error: %s", fontFamily, text, err)
	}
	defer surface.Free()

	texture, err := global.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Could create texture from font: %s", fontFamily)
	}

	_, _, width, height, _ := texture.Query()

	return &TextLabelComponent{
		Component:  ecs.NewBaseComponent(owner),
		Position:   &sdl.Rect{X: x, Y: y, W: width, H: height},
		Text:       text,
		FontFamily: fontFamily,
		Color:      color,
		Texture:    texture,
	}
}
