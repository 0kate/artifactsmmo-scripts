package dto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type FindAllMonstersResponse struct {
	Data  []Monster `json:"data"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Pages int       `json:"pages"`
}

func (f *FindAllMonstersResponse) ToPages(statusCode StatusCode) (*shared.Pages[*monsters.Monster], error) {
	switch statusCode {
	case http.StatusOK:
		monsters := make([]*monsters.Monster, len(f.Data))
		for i, data := range f.Data {
			monsters[i] = data.ToMonster()
		}
		return shared.NewPages(monsters, f.Total, f.Page, f.Size, f.Pages), nil
	case 400:
		return nil, errors.New("Monsters not found.")
	default:
		return nil, fmt.Errorf("Unexpected error for status code %d\n", statusCode)
	}
}
