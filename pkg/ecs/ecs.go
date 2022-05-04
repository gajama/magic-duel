package ecs

import (
	_ "github.com/gavmassingham/magic-duel/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
)

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
	XPos, YPos, Zindex int
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
	return "Drawable Component"
}

func (c Drawable) ID() ComponentID {
	return DRAWABLE_ID
}

const (
	SPACE_ID ComponentID = iota + 1
	SCORE_ID
	DRAWABLE_ID
	C_MAX int = iota + 1
)
