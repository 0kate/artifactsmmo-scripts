package dto

import (
	"fmt"
	"net/http"

	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type MoveData struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	// Character   Character	  `json:"character"`
}

type MoveResponse struct {
	Data MoveData `json:"data"`
}

func (r *MoveResponse) ToActionResult(statusCode StatusCode) (*characters.MoveResult, error) {
	switch statusCode {
	case http.StatusOK, 490:
		return characters.NewMoveResult(
			r.Data.Cooldown.ToActionResult(),
			r.Data.Destination.ToActionResult(),
			nil,
		), nil
	default:
		return nil, fmt.Errorf("Unexpected error with status code %d\n", statusCode)
	}
}
