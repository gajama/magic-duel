package game

import (
	"fmt"
	"strings"

	_ "github.com/gavmassingham/magic-duel/internal/config"
	"github.com/gavmassingham/magic-duel/pkg/ecs"
)

type World struct {
	Entities ecs.EntitiesMap
	LastID   ecs.Entity
}

func MakeWorld() *World {
	e := make(map[ecs.Entity][]ecs.Component)

	return &World{
		Entities: e,
		LastID:   0,
	}
}

type EntityBuilder func(ecs.Entity, *World)

func (w *World) With(comp ecs.Component) *World {
	w.Entities[w.LastID][comp.ID()] = comp
	return w
}

func (w *World) AddEntity() *World {
	ID := ecs.Entity(w.LastID + 1)
	w.Entities[ID] = make([]ecs.Component, ecs.C_MAX)
	w.LastID = ID
	return w
}

func (w *World) ListEntities() string {
	var out strings.Builder
	out.WriteString("\n\n")
	out.WriteString("Entity ID | Components\n")
	out.WriteString("----------|------------\n")
	for entity, components := range w.Entities {
		out.WriteString(fmt.Sprintf("#%08d | ", entity))
		for _, c := range components {
			if c != nil {
				out.WriteString(fmt.Sprintf("%s [ID:%v] %#v; ", c.Name(), c.ID(), c))
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}
