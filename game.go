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
)

//go:embed resources/game-bg.png
var b []byte

func init() {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	layers [][widthInTiles * heightInTiles]tile
}

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
	return nil
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}
	g.createLayers()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebiten Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
