package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

func main() {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	characterName := os.Getenv("CHARACTER_NAME")
	if characterName == "" {
		panic("CHARACTER_NAME is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	charactersActionExecutor := artifactsapi.NewCharactersActionExecutor(config)
	mapsRepository := artifactsapi.NewMapsRepository(config)

	content := shared.NewContent(shared.ContentTypeMonster, monsters.YellowSlime.String())
	pages, err := mapsRepository.FindAll(content, nil)
	if err != nil {
		panic(err)
	}

	myCharacter := characters.NewCharacter(characterName)
	targetMap := pages.Data()[0]
	position := targetMap.Position()

	result, err := charactersActionExecutor.Move(myCharacter, position)
	if err != nil {
		panic(err)
	}

	time.Sleep(result.Cooldown().RemainingSecondsInDuration() * time.Second)

	for {
		result, err := charactersActionExecutor.Fight(myCharacter)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Fight action is successfully called.\n>>> %+v\n", result)

		time.Sleep(result.Cooldown().RemainingSecondsInDuration() * time.Second)
	}
}
