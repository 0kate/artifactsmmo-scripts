package monsters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/spf13/cobra"
)

type ListOptions struct {
}

func NewListCmd() *cobra.Command {
	o := &ListOptions{}

	list := &cobra.Command{
		Use:   "list",
		Short: "List items",
		Long:  "List items",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run()
		},
	}

	return list
}

func (o *ListOptions) Run() {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
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
