package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type AssetManager struct {
	entityManager *EntityManager
	textures      map[string]*sdl.Texture
}

func (m *AssetManager) ClearData() {
	for k := range m.textures {
		delete(m.textures, k)
	}
}

func (m *AssetManager) AddTexture(textureId string, filepath string) {
	texture, err := LoadTexture(filepath)
	if err != nil {
		log.Fatalf("Asset manager failed to load texture from file: %s", filepath)
	}
	m.textures[textureId] = texture
}

func (m *AssetManager) GetTexture(textureId string) *sdl.Texture {
	return m.textures[textureId]
}
