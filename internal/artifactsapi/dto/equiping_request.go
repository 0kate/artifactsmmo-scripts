package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type EquipingRequest struct {
	Code string `json:"code"`
	Slot string `json:"slot"`
}

func NewEquipingRequest(code items.Code, slot characters.Slot) *EquipingRequest {
	return &EquipingRequest{
		Code: code.String(),
		Slot: slot.String(),
	}
}
