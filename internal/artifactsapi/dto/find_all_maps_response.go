package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/maps"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type FindAllMapsResponse struct {
	Data  []Map `json:"data"`
	Total int   `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Pages int   `json:"pages"`
}

func (r *FindAllMapsResponse) ToPages(statusCode StatusCode) (*shared.Pages[*maps.Map], error) {
	maps := make([]*maps.Map, len(r.Data))
	for i, data := range r.Data {
		maps[i] = data.ToMap()
	}

	return shared.NewPages(maps, r.Total, r.Page, r.Size, r.Pages), nil
}
