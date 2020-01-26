package core

import (
	"log"

	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadTexture(filename string) (*sdl.Texture, error) {
	surface, err := img.Load(filename)
	if err != nil {
		log.Fatalf("Could not load texture from file: %s", filename)
		return nil, err
	}
	texture, err := global.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Could create texture from surface: %s", filename)
		return nil, err
	}
	defer surface.Free()

	return texture, nil
}

func DrawTexture(texture *sdl.Texture, sourceRect *sdl.Rect, destinationRectangle *sdl.Rect, flip sdl.RendererFlip) {
	global.Renderer.CopyEx(texture, sourceRect, destinationRectangle, 0, &sdl.Point{}, flip)
}
