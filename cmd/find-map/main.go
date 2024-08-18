package main

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	mapsRepository := artifactsapi.NewMapsRepository(config)

	position := shared.NewPosition(4, -1)

	map_, err := mapsRepository.Find(position)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Map: %#v\n", map_)
	fmt.Printf("Map.Content: %#v\n", map_.Content())
}
