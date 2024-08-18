package main

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("API_TOKEN is required")
	}

	characterName := os.Getenv("CHARACTER_NAME")
	if characterName == "" {
		panic("CHARACTER_NAME is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	charactersRepository := artifactsapi.NewCharacterRepository(config)

	myCharacter := characters.NewCharacter(characterName)
	result, err := charactersRepository.Crafting(myCharacter, items.WoodenStaff, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %+v\n", result)
}
