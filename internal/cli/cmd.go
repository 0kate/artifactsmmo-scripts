package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/0kate/artifactsmmo-scripts/internal/cli/characters"
	"github.com/0kate/artifactsmmo-scripts/internal/cli/maps"
	"github.com/0kate/artifactsmmo-scripts/internal/cli/monsters"
	"github.com/spf13/cobra"
)

type ArtifactsCLICmd struct {
	cmd *cobra.Command
}

func NewArtifactsCLICmd() *ArtifactsCLICmd {
	ctx := context.Background()
	cmd := &cobra.Command{
		Use:   "artifacts-cli",
		Short: "A CLI for interacting with the Artifacts API",
		Long:  "A CLI for interacting with the Artifacts API",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			apiToken := os.Getenv("ARTIFACTS_API_TOKEN")
			if apiToken == "" {
				return fmt.Errorf("ARTIFACTS_API_TOKEN is required")
			}

			ctx = context.WithValue(ctx, "apiToken", apiToken)

			cmd.SetContext(ctx)

			return nil
		},
	}

	cmd.AddCommand(characters.NewCommand(ctx))
	cmd.AddCommand(maps.NewCommand(ctx))
	cmd.AddCommand(monsters.NewCommand(ctx))

	return &ArtifactsCLICmd{cmd: cmd}
}

func (c *ArtifactsCLICmd) Run() error {
	return c.cmd.Execute()
}
