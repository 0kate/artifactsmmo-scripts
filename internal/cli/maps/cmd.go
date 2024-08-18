package maps

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "maps",
		Short: "Interact with the maps API",
		Long:  "Interact with the maps API",
	}

	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewShowCmd())

	return cmd
}
