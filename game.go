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

var bgTiles *ebiten.Image

//go:embed resources/game-bg.png
var b []byte

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	bgTiles = ebiten.NewImageFromImage(img)
}

type Game struct{}

func (g *Game) Update() error {
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
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tiling test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
