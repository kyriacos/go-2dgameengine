package ecs_test

import (
	"testing"

	. "github.com/kyriacos/2dgameengine/ecs"
)

const FakeTestComponentOneType ComponentType = 1 << 0
const FakeTestComponentTwoType ComponentType = 1 << 1
const FakeTestComponentThreeType ComponentType = 1 << 2

type TestComponentOne struct{ Name string }

func (c *TestComponentOne) Type() ComponentType {
	return FakeTestComponentOneType
}

type TestComponentTwo struct{ Name string }

func (c *TestComponentTwo) Type() ComponentType {
	return FakeTestComponentTwoType
}

type TestComponentThree struct{ Name string }

func (c *TestComponentThree) Type() ComponentType {
	return FakeTestComponentThreeType
}

type OneSystemTest struct {
	*System
}

func (s *OneSystemTest) Update(deltaTime float64) {}
func (s *OneSystemTest) Add(e IEntity)            { s.Entities = append(s.Entities, e) }
func (s *OneSystemTest) Signature() ComponentBitMask {
	return ComponentBitMask(FakeTestComponentOneType)
}

type OneAndTwoSystemTest struct{ Entities []IEntity }

func (s *OneAndTwoSystemTest) Update(deltaTime float64) {}
func (s *OneAndTwoSystemTest) Add(e IEntity)            { s.Entities = append(s.Entities, e) }
func (s *OneAndTwoSystemTest) Remove(e IEntity) {
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
func (s *OneAndTwoSystemTest) Signature() ComponentBitMask {
	return ComponentBitMask(FakeTestComponentOneType | FakeTestComponentTwoType)
}

type ThreeSystemTest struct{ Entities []IEntity }

func (s *ThreeSystemTest) Update(deltaTime float64) {}
func (s *ThreeSystemTest) Add(e IEntity)            { s.Entities = append(s.Entities, e) }
func (s *ThreeSystemTest) Remove(e IEntity) {
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
func (s *ThreeSystemTest) Signature() ComponentBitMask {
	return ComponentBitMask(FakeTestComponentThreeType)
}

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

func TestRemoveEntity(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()
	testComp := &TestComponentOne{}

	oneSystem := &OneSystemTest{System: &System{}}
	em.Systems = append(em.Systems, oneSystem)
	em.AddEntity(e)
	em.AddComponent(e, testComp)

	em.RemoveEntity(e)
	if len(em.Entities) != 0 {
		t.Errorf("Failed to remove the entity from the entity manager!")
	}
	if len(oneSystem.Entities) != 0 {
		t.Errorf("Failed to remove the entity from the system!")
	}
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

	if oldBitMask == newBitMask ||
		em.Entities[e] == oldBitMask ||
		uint64(newBitMask)^uint64(FakeTestComponentOneType) != 0 {

		t.Errorf(`
            Failed to update the bitmask for the entity after removing a component. 
            Got:	0b%0.64b. 
            Expected:	0b%0.64b`, oldBitMask, newBitMask)
	}

}

func TestAddComponentVariadic(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	em.AddEntity(e)

	if len(em.EntityComponents[e]) > 0 {
		t.Errorf("The components should be empty!")
	}

	one := &TestComponentOne{Name: "ComponentOne"}
	two := &TestComponentTwo{Name: "ComponentTwo"}

	oldBitMask := em.GetEntityComponentBitMask(e)
	em.AddComponent(e, one, two)
	newBitMask := em.GetEntityComponentBitMask(e)

	if len(em.EntityComponents[e]) == 0 {
		t.Errorf("Entity components are empty!")
	}

	comp := em.EntityComponents[e][FakeTestComponentOneType]
	if comp != one {
		t.Errorf("Entity component for type %d is %s does not match the expected component %s",
			FakeTestComponentOneType, comp, one)
	}

	comp = em.EntityComponents[e][FakeTestComponentTwoType]
	if comp != two {
		t.Errorf("Entity component for type %d is %s does not match the expected component %s",
			FakeTestComponentTwoType, comp, two)
	}

	expectedBitMask := FakeTestComponentOneType | FakeTestComponentTwoType
	if oldBitMask == newBitMask ||
		em.Entities[e] == oldBitMask ||
		newBitMask != ComponentBitMask(expectedBitMask) {
		t.Errorf(`
            Failed to update the bitmask for the entity after adding the components. 
            Got:	0b%0.64b. 
            Expected:	0b%0.64b`, oldBitMask, newBitMask)
	}

}

func TestSystemsWithOneComponentFilter(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	oneSystem := &OneSystemTest{System: &System{}}
	em.Systems = append(em.Systems, oneSystem)

	em.AddEntity(e)
	em.AddComponent(e, &TestComponentOne{})

	if len(oneSystem.Entities) == 0 || oneSystem.Entities[0] != e {
		t.Errorf("Failed to add entity to the correct system.")
	}
}

func TestSystemsMultipleAddAndRemove(t *testing.T) {
	oneSystem := &OneSystemTest{System: &System{}}
	threeSystem := &ThreeSystemTest{}
	oneAndTwoSystem := &OneAndTwoSystemTest{}

	em := NewEntityManager()
	em.Systems = append(em.Systems, oneSystem)
	em.Systems = append(em.Systems, threeSystem)
	em.Systems = append(em.Systems, oneAndTwoSystem)

	e := NewEntity()
	em.AddEntity(e)
	one := &TestComponentOne{}
	two := &TestComponentTwo{}
	three := &TestComponentThree{}

	// add one more entity a different two components but one is different
	em.AddComponent(e, one)
	em.AddComponent(e, two)
	em.AddComponent(e, three)

	if len(oneSystem.Entities) != 1 ||
		len(oneAndTwoSystem.Entities) != 1 ||
		len(threeSystem.Entities) != 1 {

		t.Errorf("Wrong number of entities found in the systems")
	}

	em.RemoveComponent(e, one)

	if len(oneSystem.Entities) != 0 {
		// len(oneAndTwoSystem.Entities) != 0 {
		t.Errorf("Failed to remove the entity from the two systems.")
	}

	if len(threeSystem.Entities) != 1 {
		t.Errorf("Entity was removed from a system that it shouldn't have been removed from.")
	}
}

func TestSystemsWithTwoComponentsFilter(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	oneSystem := &OneSystemTest{System: &System{}}
	oneAndTwoSystem := &OneAndTwoSystemTest{}
	em.Systems = append(em.Systems, oneSystem)
	em.Systems = append(em.Systems, oneAndTwoSystem)

	em.AddEntity(e)
	one := &TestComponentOne{}
	two := &TestComponentTwo{}

	em.AddComponent(e, one)

	if len(oneSystem.Entities) != 1 || oneSystem.Entities[0] != e {
		t.Errorf("Failed to add entity to the correct system.")
	}

	em.AddComponent(e, two)

	if len(oneSystem.Entities) != 1 {
		t.Errorf("Entity was added twice to the same system!")
	}

	if len(oneAndTwoSystem.Entities) != 1 || oneAndTwoSystem.Entities[0] != e {
		t.Errorf("Failed to add entity to the second system.")
	}
}

func TestSystemsWithTwoComponentsFilterRemoval(t *testing.T) {
	em := NewEntityManager()
	e := NewEntity()

	oneSystem := &OneSystemTest{System: &System{}}
	oneAndTwoSystem := &OneAndTwoSystemTest{}
	em.Systems = append(em.Systems, oneSystem)
	em.Systems = append(em.Systems, oneAndTwoSystem)

	em.AddEntity(e)
	one := &TestComponentOne{}
	two := &TestComponentTwo{}

	em.AddComponent(e, one)

	if len(oneSystem.Entities) != 1 || oneSystem.Entities[0] != e {
		t.Errorf("Failed to add entity to the correct system.")
	}

	em.AddComponent(e, two)

	if len(oneSystem.Entities) != 1 {
		t.Errorf("Entity was added twice to the same system!")
	}

	if len(oneAndTwoSystem.Entities) != 1 || oneAndTwoSystem.Entities[0] != e {
		t.Errorf("Failed to add entity to the second system.")
	}

	em.RemoveComponent(e, one)

	if len(oneAndTwoSystem.Entities) != 0 {
		t.Errorf("Failed to remove the entity from the second system.")
	}

	if len(oneSystem.Entities) != 0 {
		t.Errorf("Failed to remove the entity from the first system.")
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
