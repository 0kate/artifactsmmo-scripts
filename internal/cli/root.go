package cli

import (
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/cli/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/cli/maps"
	"github.com/0kate/artifactsmmo-scripts/internal/cli/monsters"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	*cobra.Command
	Config *Config
}

type Config struct {
	APIToken string
}

func NewRootCmd() *RootCmd {
	root := &RootCmd{
		Command: &cobra.Command{
			Use:   "artifacts-cli",
			Short: "A CLI for interacting with the Artifacts API",
			Long:  "A CLI for interacting with the Artifacts API",
		},
		Config: &Config{},
	}

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
		if apiToken == "" {
			return fmt.Errorf("ARTIFACTS_API_TOKEN is required")
		}

		root.Config.APIToken = apiToken

		return nil
	}

	root.AddCommand(characters.NewCommand())
	root.AddCommand(maps.NewCommand())
	root.AddCommand(monsters.NewCommand())

	return root
}

func (r *RootCmd) Run() error {
	return r.Execute()
}
