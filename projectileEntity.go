package main

import (
	"github.com/kyriacos/2dgameengine/components"
	"github.com/kyriacos/2dgameengine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type ProjectileEntity struct {
	*ecs.Entity
	*components.TransformComponent
	*components.RenderComponent
}

func NewProjectileEntity() *ProjectileEntity {
	pe := &ProjectileEntity{Entity: ecs.NewEntity()}
	pe.TransformComponent = components.NewTransformComponent(0, 0, 20, 20, 32, 32, 1, pe)
	pe.RenderComponent = &components.RenderComponent{Component: ecs.NewBaseComponent(pe), Shape: sdl.Rect{}}

	return pe
}
