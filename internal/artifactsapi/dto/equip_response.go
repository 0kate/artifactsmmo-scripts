package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type EquipData struct {
	Cooldown Cooldown `json:"cooldown"`
}

type EquipResponse struct {
	Data EquipData `json:"data"`
}

func (r *EquipResponse) ToActionResult(statusCode StatusCode) (*characters.EquipResult, error) {
	return characters.NewEquipResult(characters.NewSlot(""), nil, nil), nil
}
