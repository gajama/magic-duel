package game

import (
	"fmt"
	"strings"

	_ "github.com/gavmassingham/magic-duel/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Entities map[Entity][]Component
	LastID   Entity
}

func MakeWorld() *World {
	e := make(map[Entity][]Component)

	return &World{
		Entities: e,
		LastID:   0,
	}
}

type Entity uint64

type ComponentName string

type ComponentID uint

type Component interface {
	Name() ComponentName
	ID() ComponentID
}

type Notifier interface {
	AddTransmitter() Transmit
}

type Receiver interface {
	AddReceiver() Receive
}

type EventNotify struct {
	Message string
}

type Transmit <-chan EventNotify

type Receive chan<- EventNotify

type Space struct {
	XPos, YPos int
}

func (Space) Name() ComponentName {
	return "Space Component"
}

func (c Space) ID() ComponentID {
	return SPACE_ID
}

type Score struct {
	Score int
}

func (Score) Name() ComponentName {
	return "Score Component"
}

func (c Score) ID() ComponentID {
	return SCORE_ID
}

type Drawable struct {
	Image *ebiten.Image
}

func (Drawable) Name() ComponentName {
	return "Drawable Componenet"
}

func (c Drawable) ID() ComponentID {
	return DRAWABLE_ID
}

const (
	SPACE_ID ComponentID = iota + 1
	SCORE_ID
	DRAWABLE_ID
	C_MAX int = iota
)

type EntityBuilder func(Entity, *World)

func (w *World) With(comp Component) *World {
	w.Entities[w.LastID][comp.ID()] = comp
	return w
}

func (w *World) AddEntity() *World {
	ID := Entity(w.LastID + 1)
	w.Entities[ID] = make([]Component, C_MAX)
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
				out.WriteString(fmt.Sprintf("%s [ID:%v] %#v ", c.Name(), c.ID(), c))
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}
