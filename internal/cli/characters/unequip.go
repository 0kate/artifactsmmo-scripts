package characters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type UnequipOptions struct {
	CharacterName string
}

func NewUnequipCmd(ctx context.Context) *cobra.Command {
	o := &UnequipOptions{}

	crafting := &cobra.Command{
		Use:   "unequip",
		Short: "Unequip an item from a character",
		Long:  "Unequip an item from a character",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

			myCharacter := characters.NewCharacter(o.CharacterName)

			result, err := actionExecutor.Unequip(myCharacter, characters.SlotWeapon)
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
