package ecs

import (
	_ "github.com/gavmassingham/magic-duel/internal/config"
	_ "github.com/hajimehoshi/ebiten/v2"
)

type Entity uint64

type EntitiesMap map[Entity][]Component

type NameMap map[string]Entity

func (eMap EntitiesMap) GetEntities(CIDs ...ComponentID) [][]Component {
	list := [][]Component{}
	for _, comps := range eMap {
		var add bool
		for _, CID := range CIDs {
			add = comps[CID] != nil
		}
		if add {
			list = append(list, comps)
		}
	}
	//log.Printf("%+v", list)
	return list
}
