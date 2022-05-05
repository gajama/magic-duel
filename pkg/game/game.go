package game

import (
	"fmt"
	"image"
	"strings"

	"github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	layers       [][config.WidthInTiles * config.HeightInTiles]tile
	worlds       []*World
	currentWorld uint
	nextWorldID  uint
	wrap         bool
	world        *World
}

func NewGame() *Game {

	g := &Game{}
	g.wrap = false
	g.makeWorld()
	g.createLayers()
	g.world = g.worlds[g.currentWorld]

	return g
}

func (g *Game) makeWorld() uint {
	eMap := make(map[ecs.Entity][]ecs.Component)
	nMap := make(map[string]ecs.Entity)
	ID := g.nextWorldID
	g.nextWorldID = ID + 1

	g.worlds = append(g.worlds, &World{
		Entities:       eMap,
		LastEntityID:   0,
		EntitiesByName: nMap,
		WorldID:        ID,
	})
	return ID
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
	return nil
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

	g.Render(screen)

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

func (g *Game) Render(screen *ebiten.Image) {
	for _, comps := range g.world.Entities.GetEntities(ecs.LOCATION_ID, ecs.IMAGE_ID) {
		pos := comps[ecs.LOCATION_ID].(ecs.Location)
		img := comps[ecs.IMAGE_ID].(ecs.Image).Image
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.XPos), float64(pos.YPos))
		screen.DrawImage(img, op)
	}
}

func (g *Game) AddEntity() *World {
	return g.world.addEntity()
}

func (g *Game) ListEntities() string {
	var out strings.Builder
	out.WriteString("\n\n")
	out.WriteString("Entity ID | Components\n")
	out.WriteString("----------|------------\n")
	for entity, components := range g.world.Entities {
		out.WriteString(fmt.Sprintf("#%08d | ", entity))
		for _, c := range components {
			if c != nil {
				out.WriteString(fmt.Sprintf("#:%v %#v; ", c.ID(), c))
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}
