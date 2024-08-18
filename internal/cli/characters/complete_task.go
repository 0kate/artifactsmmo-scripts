package characters

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/artifactsapi"
	"github.com/0kate/artifactsmmo-scripts/internal/characters"
	"github.com/spf13/cobra"
)

type CompleteTaskOptions struct {
	CharacterName string
}

func NewCompleteTaskCmd() *cobra.Command {
	o := &CompleteTaskOptions{}

	completeTask := &cobra.Command{
		Use:   "complete-task",
		Short: "Complete a task",
		Long:  "Complete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	completeTask.Flags().StringVarP(&o.CharacterName, "character", "c", "", "Character name")

	return completeTask
}

func (o *CompleteTaskOptions) Run() error {
	apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
	if apiToken == "" {
		panic("ARTIFACTS_API_TOKEN is required")
	}

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
}
