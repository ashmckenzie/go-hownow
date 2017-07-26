package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/ashmckenzie/go-hownow/hownow/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
}
