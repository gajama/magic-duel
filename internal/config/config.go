package config

import (
	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed resources/game-bg.png
var B []byte

//go:embed resources/king.png
var K []byte

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
