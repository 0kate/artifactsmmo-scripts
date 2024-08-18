package main

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	mapsRepository := artifactsapi.NewMapsRepository(config)

	content := shared.NewContent(shared.ContentTypeMonster, monsters.YellowSlime.String())
	pages, err := mapsRepository.FindAll(content, nil)
	if err != nil {
		panic(err)
	}

	maps := pages.Data()
	for _, m := range maps {
		fmt.Printf("Map: %+v\n", m)

		content := m.Content()
		if content != nil {
			fmt.Printf(">> Content: %+v\n", content)
		}
	}
}
