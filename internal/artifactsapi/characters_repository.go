package artifactsapi

import (
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
)

type CharactersRepository struct {
	client *Client
}

func NewCharactersRepository(config *Config) characters.CharactersRepository {
	return &CharactersRepository{
		client: NewClient(config),
	}
}

func (r *CharactersRepository) Delete(character *characters.Character) error {
	path := "/characters/delete"
	reqBody := dto.NewDeleteCharacterRequest(character.Name())

	statusCode, err := r.client.Post(path, reqBody, nil)
	if err != nil {
		return err
	}

	switch statusCode {
	case 200:
		return nil
	case 498:
		return fmt.Errorf("character %s is not found", character.Name())
	default:
		return fmt.Errorf("Unexpected status code: %d", statusCode)
	}
}
