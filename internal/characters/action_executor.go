package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type ActionExecutor interface {
	// Move character to a new position
	Move(character *Character, position *shared.Position) (*MoveResult, error)

	// Fight on the current position
	Fight(character *Character) (*FightResult, error)

	// Gathering resources on the current position
	Gathering(character *Character) (*GatheringResult, error)

	// Unequip item from a slot
	Unequip(character *Character, slot Slot) (*UnequipResult, error)

	// Craft item
	Crafting(character *Character, itemCode items.Code, quantity uint) (*CraftingResult, error)

	// Equip item to a slot
	Equip(character *Character, itemCode items.Code, slot Slot) (*EquipResult, error)

	// Complete a task
	CompleteTask(character *Character) (*CompleteTaskResult, error)
}
