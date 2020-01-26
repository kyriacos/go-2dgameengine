package core

import (
	"github.com/kyriacos/2dgameengine/core/enums"
	"github.com/kyriacos/2dgameengine/ecs"
)

type EntityManager struct {
	// entities []*ecs.Entity
	Entities      map[uint64]ecs.IEntity
	LayerEntities map[enums.LayerType][]ecs.IEntity // TODO: use a map instead with the uid
}

func (em *EntityManager) AddEntity(e ecs.IEntity, layer enums.LayerType) *ecs.IEntity {
	id := e.ID()
	em.Entities[id] = e
	em.LayerEntities[layer] = append(em.LayerEntities[layer], e)
	return &e
}

func (em *EntityManager) RemoveEntity(e ecs.IEntity, layer enums.LayerType) *ecs.IEntity {
	id := e.ID()
	em.Entities[id] = e
	// em.LayerEntities[layer] = append(em.LayerEntities[layer], e)
	// newEntities := map[enums.LayerType][]ecs.IEntity
	// for _,e:= range em.LayerEntities[layer] {

	// }
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

// destroy
// inside update() -> destroyinactiveentities()
// destroyinactive - > if !active remove from entitymanager

func (em *EntityManager) GetEntitiesByLayer(layer enums.LayerType) []ecs.IEntity {
	return em.LayerEntities[layer]
}

// update
// render
