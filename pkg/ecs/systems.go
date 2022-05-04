package ecs

import (
	"log"

	_ "github.com/gavmassingham/magic-duel/internal/config"
	_ "github.com/hajimehoshi/ebiten/v2"
)

type EntitiesMap map[Entity][]Component

type System func([]Entity)

func (eMap EntitiesMap) GetEntities(CIDs ...ComponentID) []Entity {
	list := []Entity{}
	for e, comps := range eMap {
		var add bool
		for _, CID := range CIDs {
			add = comps[CID] != nil
		}
		if add {
			list = append(list, e)
		}
	}
	log.Printf("%+v", list)
	return list
}

func (eMap EntitiesMap) Render() {
	eMap.GetEntities(SPACE_ID, DRAWABLE_ID)
}
