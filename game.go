package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tileSize     = 32
	tileSheetX   = 2
	tilesH       = 12
	tilesW       = 12
	screenWidth  = tileSize * tilesW
	screenHeight = tileSize * tilesH
)

var (
	bgTiles    *ebiten.Image
	kingSprite *ebiten.Image
)

//go:embed resources/game-bg.png
var b []byte

//go:embed resources/king.png
var k []byte

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	bgTiles = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(k))
	if err != nil {
		log.Fatal(err)
	}
	kingSprite = ebiten.NewImageFromImage(img)
}

type MapTile struct {
	PX    int
	PY    int
	Block bool
	Image image.Rectangle
}

type char struct {
	x  int
	y  int
	vx int
	vy int
}

func (c *char) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(c.x)/tileSize, float64(c.y)/tileSize)
	screen.DrawImage(kingSprite, op)
}

type Game struct {
	king *char
}

func (g *Game) Update() error {

	if g.king == nil {
		g.king = &char{x: 0, y: 0}
	}

	g.checkMove()
	g.king.x += g.king.vx
	g.king.vx = 0
	g.king.y += g.king.vy
	g.king.vy = 0
	return nil
}

func (g *Game) checkMove() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		if inBounds(g.king.y - tileSize) {
			g.king.vy -= tileSize
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		if inBounds(g.king.y + tileSize) {
			g.king.vy += tileSize
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		if inBounds(g.king.x + tileSize) {
			g.king.vx += tileSize
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		if inBounds(g.king.x - tileSize) {
			g.king.vx -= tileSize
		}

	}
}

func inBounds(v int) bool {
	if v < 0 || v > tileSize*(tilesH-1) {
		return true
	}
	return true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func makeLevel() [tilesW * tilesH]MapTile {
	var tiles [tilesW * tilesH]MapTile
	grass := image.Rect(0, 0, tileSize, tileSize)

	dirt := image.Rect(tileSize, 0, tileSize+tileSize, tileSize)

	for i := 0; i <= tilesW; i++ {
		for j := 0; j <= tilesH; j++ {

			if (i+j)%2 == 0 {
				tiles[i*tilesH+j] = MapTile{
					PX:    i * tileSize,
					PY:    j * tileSize,
					Block: false,
					Image: grass,
				}
				continue
			}
			tiles[i*tilesH+j] = MapTile{
				PX:    i * tileSize,
				PY:    j * tileSize,
				Block: false,
				Image: dirt,
			}
		}
	}
	return tiles
}

func (g *Game) Draw(screen *ebiten.Image) {
	const xNum = screenWidth / tileSize

	g.king.draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("x: %v, y:%v", g.king.x, g.king.y))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
