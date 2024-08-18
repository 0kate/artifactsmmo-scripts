package main

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	monstersRepository := artifactsapi.NewMonstersRepository(config)

	pages, err := monstersRepository.FindAll()
	if err != nil {
		panic(err)
	}

	monsters := pages.Data()
	for i, m := range monsters {
		fmt.Printf("Monster %d\n", i)
		fmt.Printf(" Name:  %+v\n", m.Name())
		fmt.Printf(" Code:  %+v\n", m.Code())
		fmt.Printf(" Level: %+v\n", m.Level())
		fmt.Printf(" HP:    %+v\n", m.HP())
	}
}
