package characters

import (
	"context"

	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "characters",
		Short: "Interact with the characters API",
		Long:  "Interact with the characters API",
	}

	cmd.AddCommand(NewMoveCmd(ctx))
	cmd.AddCommand(NewCraftingCmd(ctx))
	cmd.AddCommand(NewDeleteCmd(ctx))
	cmd.AddCommand(NewEquipCmd(ctx))
	cmd.AddCommand(NewFightCmd(ctx))
	cmd.AddCommand(NewGatheringCmd(ctx))
	cmd.AddCommand(NewUnequipCmd(ctx))
	cmd.AddCommand(NewCompleteTaskCmd(ctx))

	return cmd
}
