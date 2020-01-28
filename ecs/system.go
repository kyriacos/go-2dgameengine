package ecs

type ISystem interface {
	Update(deltaTime float64)
	Add(IEntity) // ADD(IEntity, componentsItNeeds IComponent...)
	Remove(IEntity)
	Signature() ComponentBitMask
}

type System struct {
	Entities []IEntity
}

func (s *System) Remove(e IEntity) {
	removeEntity := func(entities []IEntity, i int) []IEntity {
		entities[i] = entities[len(entities)-1]
		return entities[:len(entities)-1]
	}

	for i, entity := range s.Entities {
		if entity == e {
			s.Entities = removeEntity(s.Entities, i)
			return
		}
	}
}
