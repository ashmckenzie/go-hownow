package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/now"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app       = kingpin.New("hownow", " Time tool for determing how now should be displayed")
	noNewLine = app.Flag("no-new-line", "Don't print a newline at the end").Short('n').Bool()

	epoch = app.Flag("epoch", "Format as seconds since epoch.").Short('e').Bool()

	bod = app.Command("bod", "Beginning of day.").Alias("sod")
	bow = app.Command("bow", "Beginning of week.").Alias("sow")

	eod = app.Command("eod", "End of day.")
	eow = app.Command("eow", "End of week.")
)

func main() {
	app.Version(version)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case bod.FullCommand():
		printIt(now.BeginningOfDay())

	case bow.FullCommand():
		printIt(now.BeginningOfWeek())

	case eod.FullCommand():
		printIt(now.EndOfDay())

	case eow.FullCommand():
		printIt(now.EndOfWeek())
	}
}

func printIt(in interface{}) {
	var out string

	if *epoch {
		out = strconv.FormatInt(now.BeginningOfDay().Unix(), 10)
	} else {
		out = now.BeginningOfDay().String()
	}

	if *noNewLine {
		fmt.Print(out)
	} else {
		fmt.Println(out)
	}
}
