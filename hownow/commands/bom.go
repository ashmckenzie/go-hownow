package commands

import (
	"github.com/spf13/cobra"
)

var bomCmd = &cobra.Command{
	Use:     "bom",
	Short:   "Beginning of month",
	Aliases: []string{"som"},
	Run: func(cmd *cobra.Command, args []string) {
		display("bom")
	},
}

func init() {
	RootCmd.AddCommand(bomCmd)
}
