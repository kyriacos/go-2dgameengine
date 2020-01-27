package entities

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
)

type RadarEntity struct {
	*ecs.Entity                    `json:"Entity"`
	*components.TransformComponent `json:"TransformComponent"`
	*components.SpriteComponent    `json:"SpriteComponent"`
}

func NewRadarEntity(am *core.AssetManager) *RadarEntity {
	e := &RadarEntity{Entity: ecs.NewEntity()}
	e.TransformComponent = components.NewTransformComponent(720, 15, 0, 0, 64, 64, 1, e.Entity)
	e.SpriteComponent = components.NewAnimatedSpriteComponent(am, e.TransformComponent, "radar-image", 8, 150, false, true, e.Entity)

	return e
}

func (e *RadarEntity) RenderType() ecs.Renderable {
	return e.SpriteComponent
}

// func (e *RadarEntity) UnmarshalJSON(data []byte) error {
// 	var Key map[string][]string
// 	json.Unmarshal(data, &Key)
// 	fmt.Printf("%+v", Key)
// 	return nil
// }
