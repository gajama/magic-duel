package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize     = 32
	tileXNum     = 2
	n            = 12
	screenWidth  = tileSize * n
	screenHeight = tileSize * n
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

type char struct {
	x int
	y int
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
		g.king = &char{x: tileSize, y: tileSize}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.king.y -= tileSize
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.king.y += tileSize
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.king.x += tileSize
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.king.x -= tileSize
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	const xNum = screenWidth / tileSize
	for i := 0; i <= xNum; i++ {
		for j := 0; j <= xNum; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize), float64(j*tileSize))

			sx := (i + 1*j) % 2 * tileSize
			sy := 0
			screen.DrawImage(bgTiles.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)

		}
	}

	g.king.draw(screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
