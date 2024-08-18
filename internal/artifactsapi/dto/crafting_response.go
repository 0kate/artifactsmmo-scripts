package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type CraftingData struct {
	Cooldown Cooldown `json:"cooldown"`
}

type CraftingResponse struct {
	Data CraftingData `json:"data"`
}

func (r *CraftingResponse) ToActionResult(statusCode StatusCode) (*characters.CraftingResult, error) {
	return characters.NewCraftingResult(nil, nil, nil), nil
}
