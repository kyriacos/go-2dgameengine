package ecs

type Component struct {
	Owner IEntity `json:"-"`
}

func NewBaseComponent(o IEntity) *Component {
	return &Component{Owner: o}
}

// func (c *Component) Owner() *IEntity {
// 	return c.Owner
// }
