package integrations

import (
	"github.com/gavmassingham/magic-duel/pkg/ecs"
	"github.com/gavmassingham/magic-duel/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func Render(img *ebiten.Image) error {
	return nil
}

func RegisterEbiten(g *game.Game) {
	g.RegisterPlatform(ecs.EBITEN, Render)
}
