package main

import "github.com/kyriacos/2dgameengine/ecs"

type EntityManager struct {
	// entities []*ecs.Entity
	entities map[uint64]ecs.IEntity
}

func (em *EntityManager) AddEntity(e ecs.IEntity) *ecs.IEntity {
	id := e.ID()
	em.entities[id] = e
	return &e
}

func (em *EntityManager) GetEntity(n string) *ecs.Entity {
	// for _, e := range em.entities {
	// 	if e.Name == n {
	// 		return e
	// 	}
	// }

	return nil
}

// update
// render
