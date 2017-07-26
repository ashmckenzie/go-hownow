package commands

import (
	"github.com/spf13/cobra"
)

var eomCmd = &cobra.Command{
	Use:   "eom",
	Short: "End of month",
	Run: func(cmd *cobra.Command, args []string) {
		display("eom")
	},
}

func init() {
	RootCmd.AddCommand(eomCmd)
}
