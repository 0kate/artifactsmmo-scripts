package characters

import (
	"context"
	"fmt"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type CompleteTaskOptions struct {
	CharacterName string
}

func NewCompleteTaskCmd(ctx context.Context) *cobra.Command {
	o := &CompleteTaskOptions{}

	completeTask := &cobra.Command{
		Use:   "complete-task",
		Short: "Complete a task",
		Long:  "Complete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiToken := ctx.Value("apiToken").(string)
			config := artifactsapi.NewDefaultConfig(apiToken)
			actionExecutor := artifactsapi.NewCharactersActionExecutor(config)

			myCharacter := characters.NewCharacter(o.CharacterName)
			result, err := actionExecutor.CompleteTask(myCharacter)
			if err != nil {
				return err
			}

			fmt.Println("Task completed.")
			fmt.Printf(">> %+v\n", result)

			return nil
		},
	}

	completeTask.Flags().StringVarP(&o.CharacterName, "character", "c", "", "Character name")

	return completeTask
}
