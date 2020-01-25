package systems

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/kyriacos/2dgameengine/global"
	"github.com/veandco/go-sdl2/sdl"
)

type collisionEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.ColliderComponent
}

type CollisionSystem struct {
	entities []*collisionEntity
	AManager *core.AssetManager
	Camera   *components.CameraComponent
}

func (s *CollisionSystem) Add(e *ecs.Entity, tc *components.TransformComponent, cc *components.ColliderComponent) {
	s.entities = append(s.entities, &collisionEntity{
		Entity:             e,
		TransformComponent: tc,
		ColliderComponent:  cc,
	})
}
func (s *CollisionSystem) Update(dt float64) {
	for _, e := range s.entities {
		transform := e.TransformComponent
		colliderComponent := e.ColliderComponent

		collider, destinationRectangle := colliderComponent.Collider, colliderComponent.DestinationRectangle
		collider.X = int32(transform.Position.X)
		collider.Y = int32(transform.Position.Y)
		collider.W = int32(transform.Width * transform.Scale)
		collider.H = int32(transform.Height * transform.Scale)

		destinationRectangle.X = collider.X - int32(s.Camera.Position.X)
		destinationRectangle.Y = collider.Y - int32(s.Camera.Position.Y)

		tag := s.checkEntityCollisions(e)
		if tag == "enemy" {
			global.Running = false
		}
	}
}

func (s *CollisionSystem) checkEntityCollisions(e *collisionEntity) string {
	for _, otherEntity := range s.entities {
		if e == otherEntity {
			continue
		}
		if checkRectangleCollision(e.ColliderComponent.Collider, otherEntity.ColliderComponent.Collider) {
			return otherEntity.ColliderComponent.Tag
		}
	}
	return ""
}

func checkRectangleCollision(rectA, rectB *sdl.Rect) bool {
	return (rectA.X+rectA.W >= rectB.X &&
		rectB.X+rectB.W >= rectA.X &&
		rectA.Y+rectA.H >= rectB.Y &&
		rectB.Y+rectB.H >= rectA.Y)
}
