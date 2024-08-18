package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

type CraftingRequest struct {
	Code     string `json:"code"`
	Quantity uint   `json:"quantity"`
}

func NewCraftingRequest(code items.Code, quantity uint) *CraftingRequest {
	return &CraftingRequest{
		Code:     code.String(),
		Quantity: quantity,
	}
}
