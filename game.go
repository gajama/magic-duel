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
	"github.com/gavmassingham/magic-duel/pkg/game"
)

//go:embed resources/game-bg.png
var b []byte

//go:embed resources/king.png
var k []byte

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

}

type Game struct {
	layers     [][config.WidthInTiles * config.HeightInTiles]tile
	characters map[string]*Char
	world      *game.World
	wrap       bool
}

type Char struct {
	name  string
	game  *Game
	img   *ebiten.Image
	moved bool
	score int
}

type axis bool
type tile struct {
	ind  int
	lenX int
	lenY int
	name string
}

func (g *Game) createLayers() {
	for n := 0; n < config.NumLayers; n++ {
		var a [config.WidthInTiles * config.HeightInTiles]tile
		g.layers = append(g.layers, a)
	}
	for n := 0; n < config.WidthInTiles*config.HeightInTiles; n++ {
		g.layers[0][n] = tile{
			ind: 1,
		}
	}
}

func (g *Game) reset() {
	for n := 0; n < config.WidthInTiles*config.HeightInTiles; n++ {
		g.layers[0][n].ind = 1
	}
}

func (g *Game) Update() error {
	/* king := g.characters["king"]

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
	king.moved = false */

	return nil
}

func (c *Char) outOfBounds(a axis, move int) int {
	/* 	pos := c.yPos
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
	   	c.moved = true */
	return move
}

func keyDelayRepeat(k ebiten.Key) bool {
	if ebiten.IsKeyPressed(k) && (inpututil.KeyPressDuration(k) < 2 || inpututil.KeyPressDuration(k) > 10) {
		return true
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	const xNum = config.ScreenWidth / config.TileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*config.TileSize), float64((i/xNum)*config.TileSize))

			sx := (t.ind % config.TileXNum) * config.TileSize
			sy := (t.ind / config.TileXNum) * config.TileSize
			screen.DrawImage(config.TilesImage.SubImage(image.Rect(sx, sy, sx+config.TileSize, sy+config.TileSize)).(*ebiten.Image), op)
		}
	}
	/*
		king := g.characters["king"]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(king.xPos), float64(king.yPos))
		screen.DrawImage(king.img, op)

		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %v", king.score)) */
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	g := &Game{}
	/* 	g.wrap = false
	   	g.createLayers()
	   	g.characters = make(map[string]*Char) */
	g.world = game.MakeWorld()

	g.world.AddEntity().With(
		game.Space{XPos: 2, YPos: 2},
	)

	log.Print(g.world.ListEntities())

	/* 	ebiten.SetWindowSize(config.ScreenWidth*2, config.ScreenHeight*2)
	   	ebiten.SetWindowTitle("Game")
	   	if err := ebiten.RunGame(g); err != nil {
	   		log.Fatal(err)
	   	} */
}
