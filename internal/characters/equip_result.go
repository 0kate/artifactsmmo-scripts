package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type EquipResult struct {
	cooldown  *Cooldown
	slot      Slot
	item      *items.Item
	character *Character
}

func NewEquipResult(slot Slot, item *items.Item, character *Character) *EquipResult {
	return &EquipResult{
		slot:      slot,
		item:      item,
		character: character,
	}
}
