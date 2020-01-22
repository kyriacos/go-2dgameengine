package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/entities"
)

type GameMap struct {
	TextureID     string
	Scale         int32
	TileSize      int32
	Entities      []*entities.TileEntity
	entityManager *core.EntityManager
	assetManager  *core.AssetManager
}

func NewGameMap(em *core.EntityManager, am *core.AssetManager, textureID string, scale, tileSize int32) *GameMap {
	return &GameMap{
		entityManager: em,
		assetManager:  am,
		TextureID:     textureID,
		Scale:         scale,
		TileSize:      tileSize,
	}
}

func (m *GameMap) LoadMap(filePath string, sizeX, sizeY int) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to load map: %s. Error: %s", filePath, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	m.Entities = []*entities.TileEntity{}

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {

			// map has values such as 21,21,23
			//                        yx,yx,yx
			// 21 = tile at row 2, column 1
			rY, _, _ := r.ReadRune()
			rX, _, _ := r.ReadRune()
			r.ReadRune()

			// we want the integer value of the ASCII (UTF-8) representation of the digit
			// rune - '0' or int(rune - '0')
			sourceRectY := int32((rY - '0') * m.TileSize)
			sourceRectX := int32((rX - '0') * m.TileSize)
			positionX := int32(x) * m.Scale * m.TileSize
			positionY := int32(y) * m.Scale * m.TileSize

			m.Entities = append(m.Entities, m.AddTile(sourceRectX, sourceRectY, positionX, positionY))
		}
	}
}

func (m *GameMap) AddTile(sourceX, sourceY, x, y int32) *entities.TileEntity {
	tileEntity := entities.NewTileEntity(
		m.assetManager,
		m.TextureID,
		sourceX,
		sourceY,
		x,
		y,
		m.TileSize,
		m.Scale,
	)
	m.entityManager.AddEntity(tileEntity, core.TileMapLayer)

	return tileEntity

}
