package maps

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/monsters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
	"github.com/spf13/cobra"
)

type ListOptions struct {
}

func NewListCmd(ctx context.Context) *cobra.Command {
	list := &cobra.Command{
		Use:   "list",
		Short: "List items",
		Long:  "List items",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
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

			return nil
		},
	}

	return list
}
