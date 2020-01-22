package core

type LayerType int

const (
	TileMapLayer LayerType = iota
	VegetationLayer
	EnemyLayer
	PlayerLayer
	ProjectileLayer
	UILayer

	NumLayers int = iota
)
