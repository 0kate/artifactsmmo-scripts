package monsters

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monsters",
		Short: "Interact with the monsters API",
		Long:  "Interact with the monsters API",
	}

	cmd.AddCommand(NewListCmd())

	return cmd
}
