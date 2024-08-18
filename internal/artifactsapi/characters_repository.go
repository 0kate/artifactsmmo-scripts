package artifactsapi

import (
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type CharacterRepository struct {
	client *Client
}

func NewCharacterRepository(config *Config) *CharacterRepository {
	return &CharacterRepository{
		client: NewClient(config),
	}
}

func (r *CharacterRepository) Move(character *characters.Character, position *shared.Position) (*characters.MoveResult, error) {
	path := fmt.Sprintf("/my/%s/action/move", character.Name())
	reqBody := dto.NewMoveRequest(position)
	moveResponse := dto.MoveResponse{}

	statusCode, err := r.client.Post(path, reqBody, &moveResponse)
	if err != nil {
		return nil, err
	}

	return moveResponse.ToActionResult(statusCode)
}

func (r *CharacterRepository) Fight(character *characters.Character) (*characters.FightResult, error) {
	path := fmt.Sprintf("/my/%s/action/fight", character.Name())
	fightResponse := dto.FightResponse{}

	statusCode, err := r.client.Post(path, nil, &fightResponse)
	if err != nil {
		return nil, err
	}

	return fightResponse.ToActionResult(statusCode)
}

func (r *CharacterRepository) Gathering(character *characters.Character) (*characters.GatheringResult, error) {
	path := fmt.Sprintf("/my/%s/action/gathering", character.Name())
	gatheringResponse := dto.GatheringResponse{}

	statusCode, err := r.client.Post(path, nil, &gatheringResponse)
	if err != nil {
		return nil, err
	}

	return gatheringResponse.ToActionResult(statusCode)
}

func (r *CharacterRepository) Unequip(character *characters.Character, slot characters.Slot) (*characters.UnequipResult, error) {
	path := fmt.Sprintf("/my/%s/action/unequip", character.Name())
	reqBody := dto.NewUnequipRequest(slot)
	unequipResponse := dto.UnequipResponse{}

	statusCode, err := r.client.Post(path, reqBody, &unequipResponse)
	if err != nil {
		return nil, err
	}

	return unequipResponse.ToActionResult(statusCode)
}

func (r *CharacterRepository) Crafting(character *characters.Character, itemCode items.Code, quantity uint) (*characters.CraftingResult, error) {
	path := fmt.Sprintf("/my/%s/action/crafting", character.Name())
	reqBody := dto.NewCraftingRequest(itemCode, quantity)
	craftingResponse := dto.CraftingResponse{}

	statusCode, err := r.client.Post(path, reqBody, &craftingResponse)
	if err != nil {
		return nil, err
	}

	return craftingResponse.ToActionResult(statusCode)
}

func (r *CharacterRepository) Equip(character *characters.Character, itemCode items.Code, slot characters.Slot) (*characters.EquipResult, error) {
	path := fmt.Sprintf("/my/%s/action/equip", character.Name())
	reqBody := dto.NewEquipingRequest(itemCode, slot)
	equipResponse := dto.EquipResponse{}

	statusCode, err := r.client.Post(path, reqBody, &equipResponse)
	if err != nil {
		return nil, err
	}

	return equipResponse.ToActionResult(statusCode)
}
