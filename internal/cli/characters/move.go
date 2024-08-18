package characters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
	"github.com/spf13/cobra"
)

type MoveOptions struct {
	CharacterName string
	X             int
	Y             int
}

func NewMoveCmd() *cobra.Command {
	o := &MoveOptions{}

	move := &cobra.Command{
		Use:   "move",
		Short: "Move a character to a new position",
		Long:  "Move a character to a new position",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	move.Flags().StringVar(&o.CharacterName, "character", "", "The name of the character to move")
	move.Flags().IntVar(&o.X, "x", 0, "The x coordinate to move the character to")
	move.Flags().IntVar(&o.Y, "y", 0, "The y coordinate to move the character to")

	return move
}

func (m *MoveOptions) Run() error {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

	config := artifactsapi.NewDefaultConfig(apiToken)
	actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

	myCharacter := characters.NewCharacter(m.CharacterName)
	position := shared.NewPosition(m.X, m.Y)

	result, err := actionExecutor.Move(myCharacter, position)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %+v\n", result)
	fmt.Printf("Cooldown: %+v\n", result.Cooldown())

	return nil
}
