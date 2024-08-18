package dto

import (
	"fmt"
	"net/http"

	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type CompleteTaskResponse struct {
	Cooldown Cooldown `json:"cooldown"`
}

func (r *CompleteTaskResponse) ToActionResult(statusCode StatusCode) (*characters.CompleteTaskResult, error) {
	switch statusCode {
	case http.StatusOK:
		return characters.NewCompleteTaskResult(r.Cooldown.ToActionResult(), nil, nil), nil
	case 486:
		return nil, fmt.Errorf("An action is already in progress by your character.")
	case 487:
		return nil, fmt.Errorf("Character has no task.")
	case 488:
		return nil, fmt.Errorf("Character has not completed the task.")
	case 497:
		return nil, fmt.Errorf("Character inventory is full.")
	case 498:
		return nil, fmt.Errorf("Character not found.")
	case 499:
		return nil, fmt.Errorf("Character in cooldown.")
	case 598:
		return nil, fmt.Errorf("Tasks Master not found on this map.")
	default:
		return nil, fmt.Errorf("Unexpected error with status code: %d", statusCode)
	}
}
