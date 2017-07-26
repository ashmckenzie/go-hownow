package commands

import (
	"github.com/spf13/cobra"
)

var eodCmd = &cobra.Command{
	Use:   "eod",
	Short: "End of day",
	Run: func(cmd *cobra.Command, args []string) {
		display("eod")
	},
}

func init() {
	RootCmd.AddCommand(eodCmd)
}
