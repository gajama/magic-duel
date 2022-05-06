package game

import (
	"fmt"
	"strings"

	"github.com/gavmassingham/magic-duel/pkg/ecs"
)

var P Platform

type Platform interface {
	Runner()
	Loader()
	Label() string
}

func RegisterPlatform(p Platform) {
	P = p
}

type World struct {
	scenes       []*Scene
	currentScene uint
	nextSceneID  uint
	wrap         bool
	scene        *Scene
}

func NewGame() *World {
	g := &World{}
	g.wrap = false
	g.makeWorld()
	g.scene = g.scenes[g.currentScene]

	return g
}

func (g *World) makeWorld() uint {
	eMap := make(map[ecs.Entity][]ecs.Component)
	nMap := make(map[string]ecs.Entity)
	ID := g.nextSceneID
	g.nextSceneID = ID + 1

	w := &Scene{
		Entities:       eMap,
		LastEntityID:   0,
		EntitiesByName: nMap,
		SceneID:        ID,
	}
	w.makeLayers()
	g.scenes = append(g.scenes, w)
	return ID
}

func (g *World) GetComponentsFor(sys *ecs.System) *ecs.System {
	sys.Components = g.scene.Entities.GetEntities(sys.IDs...)
	return sys
}

func (g *World) AddEntity() *Scene {
	return g.scene.addEntity()
}

func (g *World) ListEntities() string {
	var out strings.Builder
	out.WriteString("\n\n")
	out.WriteString("Entity ID | Components\n")
	out.WriteString("----------|------------\n")
	for entity, components := range g.scene.Entities {
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
