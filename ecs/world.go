package ecs

type World struct {
	systems []System
}

func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) Update(dt float64) {
	for _, system := range w.systems {
		system.Update(dt)
	}
}
