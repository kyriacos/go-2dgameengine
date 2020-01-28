package ecs

type ComponentBitMask uint64
type ComponentType uint64

const (
	NoComponents ComponentType = iota

	NumComponentTypes uint64 = iota
)

type IComponent interface {
	Type() ComponentType
}

// type Component struct {}

// func NewBaseComponent() *Component {
// 	return &Component{}
// }
