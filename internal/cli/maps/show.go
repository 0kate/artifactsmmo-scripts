package maps

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
	"github.com/spf13/cobra"
)

type ShowOptions struct {
	X int
	Y int
}

func NewShowCmd(ctx context.Context) *cobra.Command {
	o := &ShowOptions{}

	list := &cobra.Command{
		Use:   "show",
		Short: "Show map",
		Long:  "Show map",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			mapsRepository := artifactsapi.NewMapsRepository(config)

			position := shared.NewPosition(o.X, o.Y)

			map_, err := mapsRepository.Find(position)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Map: %#v\n", map_)
			fmt.Printf("Map.Content: %#v\n", map_.Content())

			return nil
		},
	}

	list.Flags().IntVarP(&o.X, "x", "x", 0, "X coordinate")
	list.Flags().IntVarP(&o.Y, "y", "y", 0, "Y coordinate")

	return list
}
