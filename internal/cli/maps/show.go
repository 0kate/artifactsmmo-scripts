package maps

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
	"github.com/spf13/cobra"
)

type ShowOptions struct {
	X int
	Y int
}

func NewShowCmd() *cobra.Command {
	o := &ShowOptions{}

	list := &cobra.Command{
		Use:   "show",
		Short: "Show map",
		Long:  "Show map",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run()
		},
	}

	list.Flags().IntVarP(&o.X, "x", "x", 0, "X coordinate")
	list.Flags().IntVarP(&o.Y, "y", "y", 0, "Y coordinate")

	return list
}

func (o *ShowOptions) Run() {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	mapsRepository := artifactsapi.NewMapsRepository(config)

	position := shared.NewPosition(o.X, o.Y)

	map_, err := mapsRepository.Find(position)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Map: %#v\n", map_)
	fmt.Printf("Map.Content: %#v\n", map_.Content())
}
