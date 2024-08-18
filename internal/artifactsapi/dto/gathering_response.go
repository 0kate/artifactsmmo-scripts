package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type GatheringData struct {
	Cooldown Cooldown `json:"cooldown"`
}

type GatheringResponse struct {
	Data GatheringData `json:"data"`
}

func (r *GatheringResponse) ToActionResult(statusCode StatusCode) (*characters.GatheringResult, error) {
	return characters.NewGatheringResult(nil, nil), nil
}
