package commands

import (
	"fmt"
	"time"

	"github.com/ashmckenzie/go-hownow/hownow"
	"github.com/spf13/cobra"
)

var debug, showVersion, noNewLine, epoch bool
var offset int

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:           "hownow",
	Short:         "How now should be displayed",
	Long:          `How now should be displayed`,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("%s %s\n", cmd.Use, hownow.Version)
		} else {
			cmd.Help()
		}
	},
}

// Execute ...
func Execute() error {
	err := RootCmd.Execute()
	return err
}

func process(command string) string {
	h := hownow.New(time.Now(), offset)
	return h.Process(command, epoch)
}

func display(command string) {
	out := process(command)
	if !noNewLine {
		out = fmt.Sprintf("%s\n", out)
	}

	fmt.Print(out)
}

func init() {
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show version")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debugging")

	RootCmd.PersistentFlags().BoolVarP(&noNewLine, "no-new-line", "n", false, "don't print a newline at the end")
	RootCmd.PersistentFlags().BoolVarP(&epoch, "epoch", "e", false, "format as seconds since epoch")

	RootCmd.PersistentFlags().IntVar(&offset, "offset", 0, "offset now")
}
