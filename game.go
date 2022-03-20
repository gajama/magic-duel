package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 384
	screenHeight = 384
)

const (
	tileSize = 32
	tileXNum = 2
)

const (
	widthInTiles  = screenWidth / tileSize
	heightInTiles = screenHeight / tileSize
)

const (
	numLayers = 1
)

var (
	tilesImage *ebiten.Image
	kingImage  *ebiten.Image
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
	tilesImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(k))
	if err != nil {
		log.Fatal(err)
	}
	kingImage = ebiten.NewImageFromImage(img)

}

type Game struct {
	layers     [][widthInTiles * heightInTiles]tile
	characters map[int]*Char
	wrap       bool
}

type Char struct {
	name string
	game *Game
	xPos int
	yPos int
	img  *ebiten.Image
}

type axis bool
type tile struct {
	ind int
}

func (g *Game) createLayers() {
	for n := 0; n < numLayers; n++ {
		var a [widthInTiles * heightInTiles]tile
		g.layers = append(g.layers, a)
	}
	// Checkerboard effect
	for n := 0; n < screenWidth/tileSize*screenHeight/tileSize; n++ {
		if n/widthInTiles%2 == 0 {
			g.layers[0][n] = tile{
				ind: n % 2,
			}
			continue
		}
		g.layers[0][n] = tile{
			ind: -((n % 2) - 1), // swap 0 and 1
		}
	}
}

func (g *Game) Update() error {

	king := g.characters[1]

	switch {
	case keyDelayRepeat(ebiten.KeyArrowUp):
		time.Sleep(100 * time.Millisecond)
		king.yPos += king.outOfBounds(false, -tileSize)
	case keyDelayRepeat(ebiten.KeyArrowDown):
		time.Sleep(100 * time.Millisecond)
		king.yPos += king.outOfBounds(false, tileSize)
	case keyDelayRepeat(ebiten.KeyArrowLeft):
		time.Sleep(100 * time.Millisecond)
		king.xPos += king.outOfBounds(true, -tileSize)
	case keyDelayRepeat(ebiten.KeyArrowRight):
		time.Sleep(100 * time.Millisecond)
		king.xPos += king.outOfBounds(true, tileSize)
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		king.xPos = 0
		king.yPos = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		g.wrap = !g.wrap
	}

	return nil
}

func (c *Char) outOfBounds(a axis, move int) int {
	pos := c.yPos
	max := tileSize * (heightInTiles - 1)
	if a {
		pos = c.xPos
		max = tileSize * (widthInTiles - 1)
	}

	if pos+move < 0 {
		if c.game.wrap {
			return max
		}
		return 0
	}

	if pos+move > max {
		if c.game.wrap {
			return -max
		}
		return 0
	}
	return move
}

func keyDelayRepeat(k ebiten.Key) bool {
	if ebiten.IsKeyPressed(k) && (inpututil.KeyPressDuration(k) < 2 || inpututil.KeyPressDuration(k) > 10) {
		return true
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	const xNum = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t.ind % tileXNum) * tileSize
			sy := (t.ind / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	king := g.characters[1]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(king.xPos), float64(king.yPos))
	screen.DrawImage(king.img, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetMaxTPS(60)
	g := &Game{}
	g.wrap = false
	g.createLayers()
	g.characters = make(map[int]*Char)
	king := &Char{"King", g, 0, 0, kingImage}
	g.characters[1] = king

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
