package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type UnequipResult struct {
	cooldown  *Cooldown
	slot      Slot
	item      *items.Item
	character *Character
}

func NewUnequipResult(slot Slot, item *items.Item, character *Character) *UnequipResult {
	return &UnequipResult{
		slot:      slot,
		item:      item,
		character: character,
	}
}
