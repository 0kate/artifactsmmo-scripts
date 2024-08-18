package artifactsapi

import (
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi/dto"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type CharactersActionExecutor struct {
	client *Client
}

func NewCharactersActionExecutor(config *Config) characters.ActionExecutor {
	return &CharactersActionExecutor{
		client: NewClient(config),
	}
}

func (r *CharactersActionExecutor) Move(character *characters.Character, position *shared.Position) (*characters.MoveResult, error) {
	path := fmt.Sprintf("/my/%s/action/move", character.Name())
	reqBody := dto.NewMoveRequest(position)
	moveResponse := dto.MoveResponse{}

	statusCode, err := r.client.Post(path, reqBody, &moveResponse)
	if err != nil {
		return nil, err
	}

	return moveResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) Fight(character *characters.Character) (*characters.FightResult, error) {
	path := fmt.Sprintf("/my/%s/action/fight", character.Name())
	fightResponse := dto.FightResponse{}

	statusCode, err := r.client.Post(path, nil, &fightResponse)
	if err != nil {
		return nil, err
	}

	return fightResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) Gathering(character *characters.Character) (*characters.GatheringResult, error) {
	path := fmt.Sprintf("/my/%s/action/gathering", character.Name())
	gatheringResponse := dto.GatheringResponse{}

	statusCode, err := r.client.Post(path, nil, &gatheringResponse)
	if err != nil {
		return nil, err
	}

	return gatheringResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) Unequip(character *characters.Character, slot characters.Slot) (*characters.UnequipResult, error) {
	path := fmt.Sprintf("/my/%s/action/unequip", character.Name())
	reqBody := dto.NewUnequipRequest(slot)
	unequipResponse := dto.UnequipResponse{}

	statusCode, err := r.client.Post(path, reqBody, &unequipResponse)
	if err != nil {
		return nil, err
	}

	return unequipResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) Crafting(character *characters.Character, itemCode items.Code, quantity uint) (*characters.CraftingResult, error) {
	path := fmt.Sprintf("/my/%s/action/crafting", character.Name())
	reqBody := dto.NewCraftingRequest(itemCode, quantity)
	craftingResponse := dto.CraftingResponse{}

	statusCode, err := r.client.Post(path, reqBody, &craftingResponse)
	if err != nil {
		return nil, err
	}

	return craftingResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) Equip(character *characters.Character, itemCode items.Code, slot characters.Slot) (*characters.EquipResult, error) {
	path := fmt.Sprintf("/my/%s/action/equip", character.Name())
	reqBody := dto.NewEquipingRequest(itemCode, slot)
	equipResponse := dto.EquipResponse{}

	statusCode, err := r.client.Post(path, reqBody, &equipResponse)
	if err != nil {
		return nil, err
	}

	return equipResponse.ToActionResult(statusCode)
}

func (r *CharactersActionExecutor) CompleteTask(character *characters.Character) (*characters.CompleteTaskResult, error) {
	path := fmt.Sprintf("/my/%s/action/task/complete", character.Name())
	completeTaskResponse := dto.CompleteTaskResponse{}

	statusCode, err := r.client.Post(path, nil, &completeTaskResponse)
	if err != nil {
		return nil, err
	}

	return completeTaskResponse.ToActionResult(statusCode)
}
