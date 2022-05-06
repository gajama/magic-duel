package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/ecs"
	"github.com/gavmassingham/magic-duel/pkg/game"
)

//go:embed resources/game-bg.png
var b []byte

//go:embed resources/king.png
var k []byte

var g *game.Game

func init() {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	config.TilesImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(k))
	if err != nil {
		log.Fatal(err)
	}
	config.KingImage = ebiten.NewImageFromImage(img)

	g = game.NewGame()

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

func keyDelayRepeat(k ebiten.Key) bool {
	if ebiten.IsKeyPressed(k) && (inpututil.KeyPressDuration(k) < 2 || inpututil.KeyPressDuration(k) > 10) {
		return true
	}
	return false
}

func main() {

	g.AddEntity().With(ecs.Location{}).With(ecs.Image{Image: config.KingImage}).With(ecs.Controllable{Current: true}).With(ecs.Counters{}).With(ecs.Name{Is: "The King"})

	g.AddEntity()

	log.Print(g.ListEntities())

	ebiten.SetWindowSize(config.ScreenWidth*2, config.ScreenHeight*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
