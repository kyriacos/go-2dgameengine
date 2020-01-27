package systems

import (
	"fmt"

	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/core"
	"github.com/kyriacos/2dgameengine/core/enums"
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
		if tag == enums.CollisionPlayerEnemy {
			global.Running = false
			fmt.Println("You lose!")
		}
		if tag == enums.CollisionPlayerProjectile {
			global.Running = false
			fmt.Println("You lose!")
		}
		if tag == enums.CollisionPlayerLevelComplete {
			global.Running = false
			fmt.Println("You won!")
			// next level
		}
	}
}

func (s *CollisionSystem) checkEntityCollisions(e *collisionEntity) enums.CollisionType {
	for _, otherEntity := range s.entities {
		if e == otherEntity {
			continue
		}
		a := e.ColliderComponent
		b := otherEntity.ColliderComponent
		if checkRectangleCollision(a.Collider, b.Collider) {
			if a.Tag == enums.ColliderTagPlayer && b.Tag == enums.ColliderTagEnemy {
				return enums.CollisionPlayerEnemy
			}
			if a.Tag == enums.ColliderTagPlayer && b.Tag == enums.ColliderTagProjectile {
				return enums.CollisionPlayerProjectile
			}
			if a.Tag == enums.ColliderTagEnemy && b.Tag == enums.ColliderTagProjectile {
				return enums.CollisionEnemyProjectile
			}
			if a.Tag == enums.ColliderTagPlayer && b.Tag == enums.ColliderTagLevelComplete {
				return enums.CollisionPlayerLevelComplete
			}
		}
	}
	return enums.CollisionNo
}

func checkRectangleCollision(rectA, rectB *sdl.Rect) bool {
	return (rectA.X+rectA.W >= rectB.X &&
		rectB.X+rectB.W >= rectA.X &&
		rectA.Y+rectA.H >= rectB.Y &&
		rectB.Y+rectB.H >= rectA.Y)
}
