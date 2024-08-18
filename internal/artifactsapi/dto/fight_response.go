package dto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type FightResult struct {
	Result characters.FightResultType `json:"result"`
}

type FightData struct {
	Cooldown Cooldown    `json:"cooldown"`
	Fight    FightResult `json:"fight"`
}

type FightResponse struct {
	Data FightData `json:"data"`
}

func (r *FightResponse) ToActionResult(statusCode StatusCode) (*characters.FightResult, error) {
	switch statusCode {
	case http.StatusOK:
		return characters.NewFightResult(
			r.Data.Cooldown.ToActionResult(),
			nil,
			nil,
		), nil
	case 486:
		return nil, errors.New("An action is already in progress by your characters.")
	case 497:
		return nil, errors.New("Character inventory is full.")
	case 498:
		return nil, errors.New("Character not found.")
	case 499:
		return nil, errors.New("Character in cooldown.")
	case 598:
		return nil, errors.New("Monster not found on this map.")
	default:
		return nil, fmt.Errorf("Unexpected error with status code %d\n", statusCode)
	}
}
