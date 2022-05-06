package main

import (
	"log"

	"github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/ecs"
	"github.com/gavmassingham/magic-duel/pkg/game"
	_ "github.com/gavmassingham/magic-duel/pkg/game/integration/ebiten"
)

var Platform game.Platform

func init() {
	Platform = game.P
	Platform.Load()
}

func main() {

	g := game.NewGame()

	g.AddEntity().With(ecs.Location{}).With(ecs.Image{Image: config.KingImage}).With(ecs.Controllable{Current: true}).With(ecs.Counters{}).With(ecs.Name{Is: "The King"})

	g.AddEntity()

	log.Print(g.ListEntities())

	log.Println(Platform.Label())

	Platform.Run(g)

}
