package ecs

type System interface {
	Update(deltaTime float64)
}
