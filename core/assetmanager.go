package core

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type AssetManager struct {
	EntityManager *EntityManager
	Textures      map[string]*sdl.Texture
}

func (m *AssetManager) ClearData() {
	for k := range m.Textures {
		delete(m.Textures, k)
	}
}

func (m *AssetManager) AddTexture(textureId string, filepath string) {
	texture, err := LoadTexture(filepath)
	if err != nil {
		log.Fatalf("Asset manager failed to load texture from file: %s", filepath)
	}
	m.Textures[textureId] = texture
}

func (m *AssetManager) GetTexture(textureId string) *sdl.Texture {
	return m.Textures[textureId]
}
