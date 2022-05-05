package game

import (
	"fmt"

	"github.com/gavmassingham/magic-duel/pkg/ecs"
)

type axis bool
type tile struct {
	ind  int
	lenX int
	lenY int
	name string
}

type World struct {
	Entities       ecs.EntitiesMap
	LastEntityID   ecs.Entity
	EntitiesByName ecs.NameMap
	WorldID        uint
}

func (w *World) With(comp ecs.Component) *World {
	w.Entities[w.LastEntityID][comp.ID()] = comp.New()
	return w
}

func (w *World) addEntity() *World {
	ID := ecs.Entity(w.LastEntityID + 1)
	w.Entities[ID] = make([]ecs.Component, ecs.C_MAX)
	w.LastEntityID = ID
	name := fmt.Sprintf("#%08d", ID)
	return w.With(ecs.Name{Is: name})
}

type counter struct {
	entity ecs.Entity
	name   string
	value  int64
	world  *World
}

func (w *World) addCounter(e ecs.Entity, name string, v int64) *counter {
	if entry, ok := w.Entities[e]; ok {
		var comp ecs.Counters
		if comp, ok := entry[ecs.COUNTERS_ID].(ecs.Counters); ok {
			w.With(comp)
		}
		comp = entry[ecs.COUNTERS_ID].(ecs.Counters)
		if _, ok := comp.Counters[name]; !ok {
			comp.Counters[name] = v
		}
	}
	return &counter{e, name, v, w}
}

func (c *counter) Overwrite() *counter {
	c.world.Entities[c.entity][ecs.COUNTERS_ID].(ecs.Counters).Counters[c.name] = c.value
	return c
}
