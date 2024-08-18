package artifactsapi

import (
	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type MonstersRepository struct {
	client *Client
}

func NewMonstersRepository(config *Config) monsters.MonstersRepository {
	return &MonstersRepository{
		client: NewClient(config),
	}
}

func (r *MonstersRepository) FindAll() (*shared.Pages[*monsters.Monster], error) {
	path := "/monsters"
	findAllMonstersResponse := dto.FindAllMonstersResponse{}

	statusCode, err := r.client.Get(path, map[string]string{}, &findAllMonstersResponse)
	if err != nil {
		return nil, err
	}

	return findAllMonstersResponse.ToPages(statusCode)
}
