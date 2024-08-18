package characters

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "characters",
		Short: "Interact with the characters API",
		Long:  "Interact with the characters API",
	}

	cmd.AddCommand(NewMoveCmd())
	cmd.AddCommand(NewCraftingCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewEquipCmd())
	cmd.AddCommand(NewFightCmd())
	cmd.AddCommand(NewGatheringCmd())
	cmd.AddCommand(NewUnequipCmd())
	cmd.AddCommand(NewCompleteTaskCmd())

	return cmd
}
