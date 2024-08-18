package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type InventorySlot struct {
	slot     Slot
	item     items.Code
	quantity uint
}
