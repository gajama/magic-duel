package config

import (
	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 384
	ScreenHeight = 384
)

const (
	TileSize = 32
	TileXNum = 2
)

const (
	WidthInTiles  = ScreenWidth / TileSize
	HeightInTiles = ScreenHeight / TileSize
)

const (
	NumLayers = 1
)

var (
	TilesImage *ebiten.Image
	KingImage  *ebiten.Image
)
