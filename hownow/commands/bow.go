package commands

import (
	"github.com/spf13/cobra"
)

var bowCmd = &cobra.Command{
	Use:     "bow",
	Short:   "Beginning of week",
	Aliases: []string{"sow"},
	Run: func(cmd *cobra.Command, args []string) {
		display("bow")
	},
}

func init() {
	RootCmd.AddCommand(bowCmd)
}
