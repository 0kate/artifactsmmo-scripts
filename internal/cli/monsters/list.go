package monsters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
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

			return nil
		},
	}

	return list
}
