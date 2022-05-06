package ebiten

import (
	"bytes"
	"image"
	"log"

	"github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/game"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenPlatform struct {
}

func (EbitenPlatform) Render(g *game.Game) {
	/* 		for _, comps := range  {
	   		pos := comps[ecs.LOCATION_ID].(ecs.Location)
	   		img := comps[ecs.IMAGE_ID].(ecs.Image).Image
	   		op := &ebiten.DrawImageOptions{}
	   		op.GeoM.Translate(float64(pos.XPos), float64(pos.YPos))
	   		screen.DrawImage(img, op)
	   	} */
}

func (EbitenPlatform) Run(g *game.Game) {
	ebiten.SetWindowSize(config.ScreenWidth*2, config.ScreenHeight*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (EbitenPlatform) Load() {
	img, _, err := image.Decode(bytes.NewReader(config.B))
	if err != nil {
		log.Fatal(err)
	}
	config.TilesImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(config.K))
	if err != nil {
		log.Fatal(err)
	}
	config.KingImage = ebiten.NewImageFromImage(img)
}

func (EbitenPlatform) Label() string {
	return "Ebiten"
}

func init() {
	e := EbitenPlatform{}
	game.RegisterPlatform(e)
}

func keyDelayRepeat(k ebiten.Key) bool {
	if ebiten.IsKeyPressed(k) && (inpututil.KeyPressDuration(k) < 2 || inpututil.KeyPressDuration(k) > 10) {
		return true
	}
	return false
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
