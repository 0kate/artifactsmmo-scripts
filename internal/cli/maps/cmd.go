package maps

import (
	"context"

	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "maps",
		Short: "Interact with the maps API",
		Long:  "Interact with the maps API",
	}

	cmd.AddCommand(NewListCmd(ctx))
	cmd.AddCommand(NewShowCmd(ctx))

	return cmd
}
