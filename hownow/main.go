package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/now"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app       = kingpin.New("hownow", " How now should be displayed")
	noNewLine = app.Flag("no-new-line", "Don't print a newline at the end").Short('n').Bool()
	epoch     = app.Flag("epoch", "Format as seconds since epoch.").Short('e').Bool()
	offset    = app.Flag("offset", "Offset now").Int()

	bod = app.Command("bod", "Beginning of day.").Alias("sod")
	bow = app.Command("bow", "Beginning of week.").Alias("sow")

	eod = app.Command("eod", "End of day.")
	eow = app.Command("eow", "End of week.")
)

func main() {
	app.Version(version)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case bod.FullCommand():
		printIt(pointInTime().BeginningOfDay())

	case bow.FullCommand():
		printIt(pointInTime().BeginningOfWeek())

	case eod.FullCommand():
		printIt(pointInTime().EndOfDay())

	case eow.FullCommand():
		printIt(pointInTime().EndOfWeek())
	}
}

func pointInTime() *now.Now {
	t := time.Now().AddDate(0, 0, *offset)
	return now.New(t)
}

func printIt(in time.Time) {
	var out string

	if *epoch {
		out = strconv.FormatInt(in.Unix(), 10)
	} else {
		out = in.String()
	}

	if *noNewLine {
		fmt.Print(out)
	} else {
		fmt.Println(out)
	}
}
