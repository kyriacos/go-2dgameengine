package ecs_test

import (
	"testing"

	. "github.com/kyriacos/2dgameengine/ecs"
)

type SystemTest struct{}

func (s *SystemTest) Update(deltaTime float64) {}
func (s *SystemTest) Add(e IEntity)            {}
func (s *SystemTest) Remove(e IEntity)         {}
func (s *SystemTest) Signature() ComponentBitMask {
	return 1
}

func TestAddSystem(t *testing.T) {
	w := NewWorld()
	if len(w.Systems) != 0 {
		t.Errorf("New world systems is not empty!\n")
	}

	s := &SystemTest{}
	w.AddSystem(s)
	if len(w.Systems) != 1 || w.Systems[0] != s {
		t.Errorf("World does not contain the right system: %s", w.Systems)
	}
}

func TestRemoveSystem(t *testing.T) {
	w := NewWorld()

	s := &SystemTest{}
	w.AddSystem(s)
	w.RemoveSystem(s)
	if len(w.Systems) != 0 {
		t.Errorf("World does not contain the right system: %s", w.Systems)
	}
}
