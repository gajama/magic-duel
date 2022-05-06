package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ComponentLabel string

type ComponentID uint

type Component interface {
	Label() ComponentLabel
	ID() ComponentID
	New() Component
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

type Location struct {
	XPos, YPos, Zindex int
}

func (Location) Label() ComponentLabel {
	return "Locatable"
}

func (c Location) ID() ComponentID {
	return LOCATION_ID
}

func (c Location) New() Component {
	return c
}

type Counters struct {
	Counters map[string]int64
}

func (Counters) Label() ComponentLabel {
	return "Countable"
}

func (c Counters) ID() ComponentID {
	return COUNTERS_ID
}

func (c Counters) New() Component {
	c.Counters = make(map[string]int64)
	return c
}

type Image struct {
	Image *ebiten.Image
}

func (Image) Label() ComponentLabel {
	return "Drawable"
}

func (c Image) ID() ComponentID {
	return IMAGE_ID
}

func (c Image) New() Component {
	return c
}

type Controllable struct {
	Current bool
}

func (Controllable) Label() ComponentLabel {
	return "Controllable"
}

func (Controllable) ID() ComponentID {
	return CONTROLLABLE_ID
}

func (c Controllable) New() Component {
	return c
}

type Name struct {
	Is string
}

func (Name) Label() ComponentLabel {
	return "Named"
}

func (Name) ID() ComponentID {
	return NAMED_ID
}

func (c Name) New() Component {
	return c
}

const (
	LOCATION_ID ComponentID = iota + 1
	COUNTERS_ID
	IMAGE_ID
	CONTROLLABLE_ID
	NAMED_ID
	C_MAX int = iota + 1
)
