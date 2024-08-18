package characters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/spf13/cobra"
)

type CraftingOptions struct {
	CharacterName string
}

func NewCraftingCmd(ctx context.Context) *cobra.Command {
	o := &CraftingOptions{}

	crafting := &cobra.Command{
		Use:   "crafting",
		Short: "Craft an item",
		Long:  "Craft an item",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

			myCharacter := characters.NewCharacter(o.CharacterName)
			result, err := actionExecutor.Crafting(myCharacter, items.WoodenStaff, 1)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %+v\n", result)

			return nil
		},
	}

	crafting.Flags().StringVar(&o.CharacterName, "character", "c", "The name of the character to move")

	return crafting
}
