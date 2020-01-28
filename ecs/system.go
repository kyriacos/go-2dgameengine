package ecs

type ISystem interface {
	Update(deltaTime float64)
	Add(IEntity)
	Remove(IEntity)
	Signature() ComponentBitMask
}
