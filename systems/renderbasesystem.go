package systems

import "github.com/kyriacos/2dgameengine/global"

type RenderBase struct{}

func (r *RenderBase) Update(dt float64) {
	clear()

	// call all rendering systems

	// TODO: then render present
}

func clear() {
	global.Renderer.SetDrawColor(21, 21, 21, 255)
	global.Renderer.Clear()
}
