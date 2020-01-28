package ecs

type EntityManager struct {
	Entities         map[Entity]ComponentBitMask
	EntityComponents map[Entity]map[ComponentType]IComponent

	Systems []ISystem
	// SystemSignatures map[ComponentBitMask][]*System
	// or just loop through and check each Signature()
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities:         map[Entity]ComponentBitMask{},
		EntityComponents: map[Entity]map[ComponentType]IComponent{},
		Systems:          []ISystem{},
	}
}

// AddSystem
// RemoveSystem

func (em *EntityManager) AddEntity(e Entity) *Entity {
	em.Entities[e] = ComponentBitMask(NoComponents)
	em.EntityComponents[e] = map[ComponentType]IComponent{}
	return &e
}

func (em *EntityManager) RemoveEntity(e Entity) {
	delete(em.Entities, e)

	em.deleteEntityComponents(e) // remove all components and then eventually delete the entry from the map

	// for _, s := range em.Systems {
	// 	s.Remove(e)
	// }
}

func (em *EntityManager) deleteEntityComponents(e Entity) {
	// component.destroy()
	// can just call removeComponent but i think that will be slower than just going through the systems once and removing the entity
}

func (em *EntityManager) GetEntityComponentBitMask(e Entity) ComponentBitMask {
	return em.Entities[e]
}

func (em *EntityManager) AddComponent(e Entity, c IComponent) {
	currentSignature := em.Entities[e]
	newSignature := currentSignature | ComponentBitMask(c.Type())

	// add to entity components
	em.Entities[e] = newSignature
	em.EntityComponents[e][c.Type()] = c

	// loop thourgh all systems that have that signature and add the entity to them

}

func (em *EntityManager) RemoveComponent(e Entity, c IComponent) {
	currentSignature := em.Entities[e]
	newSignature := currentSignature ^ ComponentBitMask(c.Type()) // flip the bit to 0

	// loop thourgh all systems that have the current signature mask and not the new one
	// if they had the current and the new one does not apply remove the entity

	// update the signature
	em.Entities[e] = newSignature
	// remove the component
	delete(em.EntityComponents[e], c.Type())
	// call destroy on the component? if there is cleanup1
}

func (em *EntityManager) GetComponent(e Entity, c ComponentType) IComponent {
	return em.EntityComponents[e][c]
}

// destroy
// inside update() -> destroyinactiveentities()
// destroyinactive - > if !active remove from entitymanager
