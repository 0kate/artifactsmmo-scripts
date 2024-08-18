package characters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	CharacterName string
}

func NewDeleteCmd() *cobra.Command {
	o := &DeleteOptions{}

	delete := &cobra.Command{
		Use:   "delete",
		Short: "Delete a character",
		Long:  "Delete a character",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	delete.Flags().StringVar(&o.CharacterName, "character", "c", "Character name to delete")

	return delete
}

func (o *DeleteOptions) Run() error {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	charactersRepository := artifactsapi.NewCharactersRepository(config)

	character := characters.NewCharacter(o.CharacterName)

	if err := charactersRepository.Delete(character); err != nil {
		return err
	}

	fmt.Println("Character is successfully deleted.")

	return nil
}
