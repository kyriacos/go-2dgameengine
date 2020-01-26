package core

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type AssetManager struct {
	EntityManager *EntityManager
	Textures      map[string]*sdl.Texture
	Fonts         map[string]*ttf.Font
}

func (m *AssetManager) ClearData() {
	for k := range m.Textures {
		m.Textures[k].Destroy()
		delete(m.Textures, k)
	}
	for k := range m.Fonts {
		m.Fonts[k].Close()
		delete(m.Fonts, k)
	}
}

func (m *AssetManager) AddTexture(textureID string, filepath string) {
	texture, err := LoadTexture(filepath)
	if err != nil {
		log.Fatalf("Asset manager failed to load texture from file: %s", filepath)
	}
	m.Textures[textureID] = texture
}

func (m *AssetManager) AddFont(fontID string, filepath string, fontSize int) {
	font, err := LoadFont(filepath, fontSize)
	if err != nil {
		log.Fatalf("Asset manager failed to load font from file: %s", filepath)
	}
	m.Fonts[fontID] = font
}

func (m *AssetManager) GetTexture(textureID string) *sdl.Texture {
	return m.Textures[textureID]
}

func (m *AssetManager) GetFont(fontID string) *ttf.Font {
	return m.Fonts[fontID]
}
