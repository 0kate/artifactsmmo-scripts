package dto

import (
	"github.com/0kate/artifactsmmo-scripts/internal/maps"
)

type FindMapResponse struct {
	Data Map `json:"data"`
}

func (r *FindMapResponse) ToMap(statusCode StatusCode) (*maps.Map, error) {
	return r.Data.ToMap(), nil
}
