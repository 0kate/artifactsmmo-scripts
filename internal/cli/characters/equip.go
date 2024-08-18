package characters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/spf13/cobra"
)

type EquipOptions struct {
	CharacterName string
}

func NewEquipCmd(ctx context.Context) *cobra.Command {
	o := &EquipOptions{}

	equip := &cobra.Command{
		Use:   "equip",
		Short: "Equip an item",
		Long:  "Equip an item",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

			myCharacter := characters.NewCharacter(o.CharacterName)
			result, err := actionExecutor.Equip(myCharacter, items.WoodenStaff, characters.SlotWeapon)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %+v\n", result)

			return nil
		},
	}

	return equip
}
