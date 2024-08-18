package artifactsapi

import (
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
	"github.com/0kate/artifactsmmo-scripts/internal/maps"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MapsRepository struct {
	client *Client
}

func NewMapsRepository(config *Config) *MapsRepository {
	return &MapsRepository{
		client: NewClient(config),
	}
}

func (r *MapsRepository) FindAll(content *shared.Content, pages *shared.Pages[*maps.Map]) (*shared.Pages[*maps.Map], error) {
	path := "/maps"
	query := map[string]string{}
	findAllMapsResponse := dto.FindAllMapsResponse{}

	if content != nil {
		query["content_type"] = content.Type()
		query["content_code"] = content.Code()
	}

	if pages != nil {
		query["page"] = string(pages.Page())
		query["size"] = string(pages.Size())
	}

	statusCode, err := r.client.Get(path, query, &findAllMapsResponse)
	if err != nil {
		return nil, err
	}

	return findAllMapsResponse.ToPages(statusCode)
}

func (r *MapsRepository) Find(position *shared.Position) (*maps.Map, error) {
	path := fmt.Sprintf("/maps/%d/%d", position.X(), position.Y())
	findMapResponse := dto.FindMapResponse{}

	statusCode, err := r.client.Get(path, map[string]string{}, &findMapResponse)
	if err != nil {
		return nil, err
	}

	return findMapResponse.ToMap(statusCode)
}
