package commands

import (
	"github.com/spf13/cobra"
)

var bodCmd = &cobra.Command{
	Use:     "bod",
	Short:   "Beginning of day",
	Aliases: []string{"sod"},
	Run: func(cmd *cobra.Command, args []string) {
		display("bod")
	},
}

func init() {
	RootCmd.AddCommand(bodCmd)
}
