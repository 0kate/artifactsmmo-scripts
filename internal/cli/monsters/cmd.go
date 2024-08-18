package monsters

import (
	"context"

	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monsters",
		Short: "Interact with the monsters API",
		Long:  "Interact with the monsters API",
	}

	cmd.AddCommand(NewListCmd(ctx))

	return cmd
}
