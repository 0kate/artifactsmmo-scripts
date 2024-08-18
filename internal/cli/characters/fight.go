package characters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type FightOptions struct {
	CharacterName string
}

func NewFightCmd() *cobra.Command {
	o := &FightOptions{}

	crafting := &cobra.Command{
		Use:   "fight",
		Short: "Fight an enemy",
		Long:  "Fight an enemy",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	crafting.Flags().StringVar(&o.CharacterName, "character", "c", "The name of the character to move")

	return crafting
}

func (o *FightOptions) Run() error {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

	myCharacter := characters.NewCharacter(o.CharacterName)
	result, err := actionExecutor.Fight(myCharacter)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %+v\n", result)

	return nil
}
