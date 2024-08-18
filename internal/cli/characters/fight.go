package characters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type FightOptions struct {
	CharacterName string
}

func NewFightCmd(ctx context.Context) *cobra.Command {
	o := &FightOptions{}

	crafting := &cobra.Command{
		Use:   "fight",
		Short: "Fight an enemy",
		Long:  "Fight an enemy",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

			myCharacter := characters.NewCharacter(o.CharacterName)
			result, err := actionExecutor.Fight(myCharacter)
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
