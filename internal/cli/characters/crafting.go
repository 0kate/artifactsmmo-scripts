package characters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/spf13/cobra"
)

type CraftingOptions struct {
	CharacterName string
}

func NewCraftingCmd() *cobra.Command {
	o := &CraftingOptions{}

	crafting := &cobra.Command{
		Use:   "crafting",
		Short: "Craft an item",
		Long:  "Craft an item",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	crafting.Flags().StringVar(&o.CharacterName, "character", "c", "The name of the character to move")

	return crafting
}

func (o *CraftingOptions) Run() error {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

	myCharacter := characters.NewCharacter(o.CharacterName)
	result, err := actionExecutor.Crafting(myCharacter, items.WoodenStaff, 1)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %+v\n", result)

	return nil
}
