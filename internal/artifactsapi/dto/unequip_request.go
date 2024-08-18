package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type UnequipRequest struct {
	Slot string `json:"slot"`
}

func NewUnequipRequest(slot characters.Slot) *UnequipRequest {
	return &UnequipRequest{
		Slot: slot.String(),
	}
}
