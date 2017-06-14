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

	bod = app.Command("bod", "Beginning of day.")
	eod = app.Command("eod", "End of day.")
)

func main() {
	app.Version(version)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case bod.FullCommand():
		printIt(beginningOfDay())

	case eod.FullCommand():
		printIt(endOfDay())
	}
}

func beginningOfDay() string {
	if *epoch {
		return strconv.FormatInt(now.BeginningOfDay().Unix(), 10)
	} else {
		return now.BeginningOfDay().String()
	}
}

func endOfDay() string {
	if *epoch {
		return strconv.FormatInt(now.EndOfDay().Unix(), 10)
	} else {
		return now.EndOfDay().String()
	}
}

func printIt(in string) {
	if *noNewLine {
		fmt.Print(in)
	} else {
		fmt.Println(in)
	}
}
