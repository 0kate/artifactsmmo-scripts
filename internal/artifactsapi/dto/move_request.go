package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MoveRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewMoveRequest(position *shared.Position) *MoveRequest {
	return &MoveRequest{
		X: position.X(),
		Y: position.Y(),
	}
}
