package systems

import (
	"fmt"
	"log"

	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type RenderLayersSystem struct {
	// entities []Renderable
	EM     *core.EntityManager
	Camera *components.CameraComponent
}

// func (r *RenderLayersSystem) Add(e Renderable) {
// 	r.entities = append(r.entities, &e)
// }

func (r *RenderLayersSystem) Update(dt float64) {
	var layerNumber core.LayerType = 0
	for layerNumber = 0; int(layerNumber) < core.NumLayers; layerNumber++ {
		entities := r.EM.GetEntitiesByLayer(layerNumber)
		for _, e := range entities {
			switch layerNumber {
			case core.TileMapLayer, core.VegetationLayer:
				render(e, r.Camera)
			case core.EnemyLayer, core.PlayerLayer:
				render(e, r.Camera)
			case core.UILayer:
				render(e, r.Camera)
			default:
				fmt.Println("not implemented yet")
			}
		}
	}
}

func render(e ecs.IEntity, camera *components.CameraComponent) {
	switch s := e.RenderType().(type) {
	case *components.TileComponent:
		s.DestinationRectangle.X = int32(s.Position.X - camera.Position.X)
		s.DestinationRectangle.Y = int32(s.Position.Y - camera.Position.Y)
		core.DrawTexture(s.Texture, s.SourceRectangle, s.DestinationRectangle, s.TileFlip)
	case *components.SpriteComponent:
		if s.IsAnimated {
			s.SourceRectangle.X = s.SourceRectangle.W * int32((int32(sdl.GetTicks())/s.AnimationSpeed)%s.NumFrames)
			s.SourceRectangle.Y = int32(s.AnimationIndex * s.TransformComponent.Height)
		}

		x := s.TransformComponent.Position.X
		y := s.TransformComponent.Position.Y

		// If the sprite is fixed then dont change its position relative to the camera
		if !s.IsFixed {
			x -= camera.Position.X
			y -= camera.Position.Y
		}

		s.DestinationRectangle.X = int32(x)
		s.DestinationRectangle.Y = int32(y)
		s.DestinationRectangle.W = int32(s.TransformComponent.Width * s.TransformComponent.Scale)
		s.DestinationRectangle.H = int32(s.TransformComponent.Height * s.TransformComponent.Scale)

		core.DrawTexture(s.Texture, s.SourceRectangle, s.DestinationRectangle, s.SpriteFlip)
	default:
		log.Fatalf("Trying to render an entity that has no RenderType defined: %s", e)
	}
}
