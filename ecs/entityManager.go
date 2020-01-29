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
// RefreshSystems (pre-load) - maybe we initialize entities with Unmarshal and then just add to the systems in one go.

func (em *EntityManager) AddEntity(e Entity) *Entity {
	em.Entities[e] = ComponentBitMask(NoComponents)
	em.EntityComponents[e] = map[ComponentType]IComponent{}
	return &e
}

func (em *EntityManager) RemoveEntity(e Entity) {

	// em.deleteEntityComponents(e) // remove all components and then eventually delete the entry from the map
	// perform clean up or something for SDL or other libraries that need it.

	for _, s := range em.Systems {
		if em.Entities[e]&s.Signature() == s.Signature() {
			s.Remove(e)
		}
	}

	delete(em.Entities, e)
}

func (em *EntityManager) GetEntityComponentBitMask(e Entity) ComponentBitMask {
	return em.Entities[e]
}

func (em *EntityManager) AddComponent(e Entity, components ...IComponent) {
	oldSignature := em.Entities[e]
	newSignature := oldSignature

	for _, c := range components {
		newSignature |= ComponentBitMask(c.Type())
		// add to entity components
		em.EntityComponents[e][c.Type()] = c
	}
	em.Entities[e] = newSignature

	// loop through all systems that have that signature and add the entity to them
	// if they dont include the entity already
	for _, s := range em.Systems {
		if oldSignature&s.Signature() != s.Signature() && newSignature&s.Signature() == s.Signature() {
			s.Add(e) // add the entity or the components the system expects? or both?
		}
	}

}

func (em *EntityManager) RemoveComponent(e Entity, components ...IComponent) {
	oldSignature := em.Entities[e]
	newSignature := oldSignature

	for _, c := range components {
		newSignature ^= ComponentBitMask(c.Type()) // flip the bit to 0
		// remove from entity components
		delete(em.EntityComponents[e], c.Type())

		// call destroy on the component? if there is cleanup1
	}
	em.Entities[e] = newSignature

	// loop through all systems that have the current signature mask and not the new one
	// if they had the current and the new one does not apply remove the entity
	for _, s := range em.Systems {
		if oldSignature&s.Signature() == s.Signature() && newSignature&s.Signature() != s.Signature() {
			s.Remove(e) // remove the entity from the system
		}
	}

}

func (em *EntityManager) GetComponent(e Entity, c ComponentType) IComponent {
	return em.EntityComponents[e][c]
}

// destroy
// inside update() -> destroyinactiveentities()
// destroyinactive - > if !active remove from entitymanager
