package ecs

import (
	"sync/atomic"
)

var _id uint64

type Renderable interface{}

type IEntity interface {
	ID() uint64
	RenderType() Renderable
}

type Entity struct {
	id uint64
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) RenderType() Renderable {
	return nil
}

// func (e *Entity) GetComponent(c *Component)

func NewEntity() *Entity {
	return &Entity{id: atomic.AddUint64(&_id, 1)}
}

// // Add Component to Entity's component list
// func (e *Entity) AddComponent(c *Component) {
// 	e.components = append(e.components, c)
// }
