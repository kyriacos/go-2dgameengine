package ecs

import (
	"sync/atomic"
)

var _id uint64

type Renderable interface{}

type IEntity interface {
	// ID() uint64
	// RenderType() Renderable
}

type Entity uint64

// func (e *Entity) ID() uint64 {
// 	return e
// }

// func (e *Entity) RenderType() Renderable {
// 	return nil
// }

func NewEntity() Entity {
	// return &Entity{id: atomic.AddUint64(&_id, 1)}
	return Entity(atomic.AddUint64(&_id, 1))
}
