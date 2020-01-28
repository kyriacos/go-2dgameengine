package ecs

type World struct {
	Systems []ISystem
}

func NewWorld() *World {
	return &World{Systems: []ISystem{}}
}

func (w *World) AddSystem(s ISystem) {
	w.Systems = append(w.Systems, s)
}

// TODO: clear the data in the system? or do we remove all entities?
// if we are storing the signature type somewhere we should remove that at least
func (w *World) RemoveSystem(s ISystem) {
	for i, sys := range w.Systems {
		if sys == s {
			w.Systems = removeSystem(w.Systems, i)
			return
		}
	}
}

func (w *World) Update(dt float64) {
	for _, system := range w.Systems {
		system.Update(dt)
	}
}

func removeSystem(s []ISystem, i int) []ISystem {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
