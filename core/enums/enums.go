package enums

type LayerType int

const (
	TileMapLayer LayerType = iota
	VegetationLayer
	EnemyLayer
	PlayerLayer
	ProjectileLayer
	ObstacleLayer
	UILayer

	NumLayers int = iota
)

type ColliderTag int

const (
	ColliderTagPlayer ColliderTag = iota
	ColliderTagEnemy
	ColliderTagVegetation
	ColliderTagProjectile
	ColliderTagLevelComplete

	NumColliderTags int = iota
)

type CollisionType int

const (
	CollisionNo CollisionType = iota
	CollisionPlayerEnemy
	CollisionPlayerProjectile
	CollisionEnemyProjectile
	CollisionPlayerVegetation
	CollisionPlayerLevelComplete

	NumCollisionTypes int = iota
)
