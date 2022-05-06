package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	EBITEN PlatformLabel = "Ebiten"
)

type PlatformLabel string

type Platform struct {
	Label  PlatformLabel
	Render func(r any) error
}

func RegisterPlatform[R RenderType](label PlatformLabel, target R) {}

type RenderFunc[R RenderType] func(R) error
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

type RenderType interface {
	*ebiten.Image
}
