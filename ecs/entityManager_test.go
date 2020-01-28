package ecs_test

import (
	"testing"

	. "github.com/kyriacos/2dgameengine/ecs"
)

func TestNewEntityManager(t *testing.T) {
	em := NewEntityManager()
	if em.Systems == nil || em.Entities == nil || em.EntityComponents == nil {
		t.Errorf("Failed to initialize Entity Manager")
	}
}
func TestAddEntity(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	em.AddEntity(e)

	if len(em.Entities) == 0 {
		t.Errorf("Failed to add entity. Entities are empty!")
	}

	if em.EntityComponents[e] == nil {
		t.Errorf("Did not initialize the components for the new entity!")
	}

	if em.Entities[e] != ComponentBitMask(NoComponents) {
		t.Errorf(
			"Something went wrong when adding the entity in the map. The bit mask is not correct. Got: 0b%0.64b. Expected: 0b%0.64b",
			em.Entities[e], ComponentBitMask(NoComponents))
	}
}

const FakeTestComponentOneType ComponentType = 1
const FakeTestComponentTwoType ComponentType = 2

type TestComponentOne struct{}

func (c *TestComponentOne) Type() ComponentType {
	return FakeTestComponentOneType
}

type TestComponentTwo struct{}

func (c *TestComponentTwo) Type() ComponentType {
	return FakeTestComponentTwoType
}

func TestAddComponent(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	em.AddEntity(e)
	if len(em.EntityComponents[e]) > 0 {
		t.Errorf("The components should be empty!")
	}

	testComp := &TestComponentOne{}
	oldBitMask := em.GetEntityComponentBitMask(e)
	em.AddComponent(e, testComp)
	newBitMask := em.GetEntityComponentBitMask(e)

	if len(em.EntityComponents[e]) == 0 || em.EntityComponents[e][FakeTestComponentOneType] != testComp {
		t.Errorf("Failed to add component for entity")
	}

	if oldBitMask == newBitMask || em.Entities[e] == oldBitMask || uint64(newBitMask)^uint64(FakeTestComponentOneType) != 0 {
		t.Errorf(`
			Failed to update the bitmask for the entity after removing a component. 
			Got: 0b%0.64b. 
			Expected: 0b%0.64b`, oldBitMask, newBitMask)
	}
}

func TestAddTwoComponents(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	em.AddEntity(e)
	if len(em.EntityComponents[e]) > 0 {
		t.Errorf("The components should be empty!")
	}

	one := &TestComponentOne{}
	two := &TestComponentTwo{}

	// ADD FIRST COMPONENT
	oldBitMask := em.GetEntityComponentBitMask(e)
	em.AddComponent(e, one)
	newBitMask := em.GetEntityComponentBitMask(e)

	if len(em.EntityComponents[e]) == 0 ||
		em.EntityComponents[e][FakeTestComponentOneType] != one {
		t.Errorf("Failed to add component for entity")
	}

	if oldBitMask == newBitMask ||
		em.Entities[e] == oldBitMask ||
		uint64(newBitMask)^uint64(FakeTestComponentOneType) != 0 {
		t.Errorf(`
			Failed to update the bitmask for the entity after adding a component. 
			Got: 0b%0.64b. 
			Expected: 0b%0.64b`, oldBitMask, newBitMask)
	}

	// ADD SECOND COMPONENT
	oldBitMask = em.GetEntityComponentBitMask(e)
	em.AddComponent(e, two)
	newBitMask = em.GetEntityComponentBitMask(e)

	if len(em.EntityComponents[e]) == 1 ||
		em.EntityComponents[e][FakeTestComponentOneType] != one ||
		em.EntityComponents[e][FakeTestComponentTwoType] != two {
		t.Errorf("Failed to add second component for entity")
	}

	if oldBitMask == newBitMask ||
		em.Entities[e] == oldBitMask ||
		uint64(newBitMask)^uint64(FakeTestComponentOneType)^uint64(FakeTestComponentTwoType) != 0 {

		t.Errorf(`
			Failed to update the bitmask for the entity after adding a second component. 
			Got: 0b%0.64b. 
			Expected: 0b%0.64b`, oldBitMask, newBitMask)
	}
}
func TestRemoveComponent(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()
	testComp := &TestComponentOne{}

	em.AddEntity(e)
	em.AddComponent(e, testComp)

	oldBitMask := em.GetEntityComponentBitMask(e)
	em.RemoveComponent(e, testComp)
	newBitMask := em.GetEntityComponentBitMask(e)

	if len(em.EntityComponents[e]) != 0 || em.EntityComponents[e][FakeTestComponentOneType] != nil {
		t.Errorf("Failed to remove component for entity")
	}

	if oldBitMask == newBitMask || em.Entities[e] == oldBitMask || newBitMask != 0 {
		t.Errorf("Failed to update the bitmask for the entity after removing a component. Got: 0b%0.64b. Expected: 0b%0.64b", newBitMask, oldBitMask)
	}
}

func TestGetComponent(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	em.AddEntity(e)

	testComp := &TestComponentOne{}
	em.AddComponent(e, testComp)

	if em.GetComponent(e, FakeTestComponentOneType) == nil {
		t.Errorf("Component does not exist")
	}

	if em.GetComponent(e, testComp.Type()) != testComp {
		t.Errorf("WTH?!")
	}
}
