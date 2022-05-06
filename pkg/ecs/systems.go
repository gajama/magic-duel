package ecs

type System struct {
	Label      string
	IDs        []ComponentID
	Components [][]Component
}

type SystemRunner interface {
	Run() error
	System
}

type Renderer interface {
	Render(dest any) error
}
