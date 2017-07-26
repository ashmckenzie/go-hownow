package commands

import (
	"github.com/spf13/cobra"
)

var eowCmd = &cobra.Command{
	Use:   "eow",
	Short: "End of week",
	Run: func(cmd *cobra.Command, args []string) {
		display("eow")
	},
}

func init() {
	RootCmd.AddCommand(eowCmd)
}
