package systems

// type E interface { }

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

// i think i would prefer to use an Interface instead of defining all this shit every type especially for the add method
// type components struct { }
type renderEntity struct {
	*ecs.Entity
	*components.RenderComponent
	*components.TransformComponent
}
type RenderSystem struct {
	entities []*renderEntity // for now we just keep the list here nice and easy.
}

// WIP: for now just push the entity we want. Eventually this should become entity and the Components
// we want this system to operate on. Or we create a new type of entity which basically is the base entity plus its components as a struct?
// Rather than have a list of components. Unless we actually need to iterate on the list of components which we dont really need to do.
func (r *RenderSystem) Add(e *ecs.Entity, rc *components.RenderComponent, tc *components.TransformComponent) {
	r.entities = append(r.entities, &renderEntity{Entity: e, RenderComponent: rc, TransformComponent: tc})
}

// needs to work on transformcomponent and rendercomponent perhaps
func (r *RenderSystem) Update(dt float64) {
	// clear() // clear the buffer
	for _, e := range r.entities {
		// renderComponent := e.RenderComponent
		t := e.TransformComponent
		drawBox(int32(t.Position.X), int32(t.Position.Y), t.Width, t.Height)
		// switch renderComponent.Shape.(type) {
		// case sdl.Rect:
		// 	drawRect()
		// default:
		// 	log.Fatal("unknown shape cannot render in RenderSystem")
		// }

	}
}

func drawBox(x, y, w, h int32) {
	rect := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: int32(w),
		H: int32(h),
	}
	global.Renderer.SetDrawColor(255, 255, 255, 255)
	global.Renderer.FillRect(rect)
}
