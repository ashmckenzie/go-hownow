package commands

import (
	"errors"
	"fmt"
	"time"

	"github.com/ashmckenzie/go-hownow/hownow"
	"github.com/spf13/cobra"
)

var debug, showVersion, noNewLine, epoch bool
var format string
var offset int

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:           "hownow",
	Short:         "How now should be displayed",
	Long:          `How now should be displayed`,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFormat()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(hownow.Version)
		}
		cmd.Help()
	},
}

// Execute ...
func Execute() error {
	err := RootCmd.Execute()
	return err
}

func process(command string) string {
	p := hownow.New(time.Now(), offset)
	t := p.Process(command)
	return p.Format(t, format, epoch)
}

func display(command string) {
	out := process(command)
	if !noNewLine {
		out = fmt.Sprintf("%s\n", out)
	}

	fmt.Print(out)
}

func validateFormat() error {
	if len(format) > 0 {
		if epoch {
			return errors.New("--format and --epoch are mutually exclusive")
		}
	} else {
		format = hownow.DefaultTimeFormat
	}

	return nil
}

func init() {
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show version")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debugging")

	RootCmd.PersistentFlags().BoolVarP(&noNewLine, "no-new-line", "n", false, "don't print a newline at the end")
	RootCmd.PersistentFlags().BoolVarP(&epoch, "epoch", "e", false, "format as seconds since epoch")

	RootCmd.PersistentFlags().StringVar(&format, "format", "", "time format, see https://golang.org/src/time/format.go")

	RootCmd.PersistentFlags().IntVar(&offset, "offset", 0, "offset now")
}
