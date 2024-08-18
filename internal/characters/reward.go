package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type Reward struct {
	itemCode items.Code
	quantity uint
}

func NewReward(itemCode items.Code, quantity uint) *Reward {
	return &Reward{
		itemCode: itemCode,
		quantity: quantity,
	}
}
