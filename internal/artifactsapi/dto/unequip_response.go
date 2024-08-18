package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type UnequipData struct {
	Cooldown Cooldown `json:"cooldown"`
}

type UnequipResponse struct {
	Data UnequipData `json:"data"`
}

func (r *UnequipResponse) ToActionResult(statusCode StatusCode) (*characters.UnequipResult, error) {
	return characters.NewUnequipResult(characters.NewSlot(""), nil, nil), nil
}
