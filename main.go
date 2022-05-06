package main

import (
	"log"

	"github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/ecs"
	"github.com/gavmassingham/magic-duel/pkg/game"
	_ "github.com/gavmassingham/magic-duel/pkg/game/integration/ebiten"
)

var Platform game.Platform

func init() {
	Platform = game.P
	Platform.Load()
}

/* func (g *Game) Update() error {
	king := g.characters["king"]

	x, y := king.xPos/config.TileSize, king.yPos/config.TileSize

	switch {
	case keyDelayRepeat(ebiten.KeyArrowUp):
		time.Sleep(100 * time.Millisecond)
		king.yPos += king.outOfBounds(false, -config.TileSize)
	case keyDelayRepeat(ebiten.KeyArrowDown):
		time.Sleep(100 * time.Millisecond)
		king.yPos += king.outOfBounds(false, config.TileSize)
	case keyDelayRepeat(ebiten.KeyArrowLeft):
		time.Sleep(100 * time.Millisecond)
		king.xPos += king.outOfBounds(true, -config.TileSize)
	case keyDelayRepeat(ebiten.KeyArrowRight):
		time.Sleep(100 * time.Millisecond)
		king.xPos += king.outOfBounds(true, config.TileSize)
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		king.xPos = 0
		king.yPos = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		g.wrap = !g.wrap
	case inpututil.IsKeyJustPressed(ebiten.KeyR):
		g.reset()
	}

	if king.moved && g.layers[0][y*config.WidthInTiles+x].ind == 1 {
		g.layers[0][y*config.WidthInTiles+x].ind = 0
		king.score += 1
	}
	king.moved = false

	return nil
} */

/* func (c *Char) outOfBounds(a axis, move int) int {
	pos := c.yPos
	max := config.TileSize * (config.HeightInTiles - 1)
	if a {
		pos = c.xPos
		max = config.TileSize * (config.WidthInTiles - 1)
	}

	if pos+move < 0 {
		if c.game.wrap {
			c.moved = true
			return max
		}
		return 0
	}

	if pos+move > max {
		if c.game.wrap {
			c.moved = true
			return -max
		}
		return 0
	}
	c.moved = true
	return move
} */

func main() {

	g := game.NewGame()

	g.AddEntity().With(ecs.Location{}).With(ecs.Image{Image: config.KingImage}).With(ecs.Controllable{Current: true}).With(ecs.Counters{}).With(ecs.Name{Is: "The King"})

	g.AddEntity()

	log.Print(g.ListEntities())

	log.Println(Platform.Label())

	Platform.Run(g)

}
