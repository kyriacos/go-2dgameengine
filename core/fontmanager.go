package core

import (
	"log"

	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func LoadFont(filename string, fontSize int) (*ttf.Font, error) {
	font, err := ttf.OpenFont(filename, fontSize)
	if err != nil {
		log.Fatalf("Could not load texture from file: %s", filename)
		return nil, err
	}

	return font, nil
}

func DrawFont(texture *sdl.Texture, position *sdl.Rect) {
	global.Renderer.Copy(texture, nil, position)
}
