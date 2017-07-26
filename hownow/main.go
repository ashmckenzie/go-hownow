package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app       = kingpin.New("hownow", " How now should be displayed")
	noNewLine = app.Flag("no-new-line", "Don't print a newline at the end").Short('n').Bool()
	epoch     = app.Flag("epoch", "Format as seconds since epoch").Short('e').Bool()
	offset    = app.Flag("offset", "Offset now").Int()

	free     = app.Command("free", "Free form")
	freeText = free.Arg("text", "Free form text").Required().String()

	bod = app.Command("bod", "Beginning of day").Alias("sod")
	eod = app.Command("eod", "End of day")

	bow = app.Command("bow", "Beginning of week").Alias("sow")
	eow = app.Command("eow", "End of week")

	bom = app.Command("bom", "Beginning of month").Alias("som")
	eom = app.Command("eom", "End of month")
)

func main() {
	fmt.Printf(parse(os.Args[1:]))
}

func whenTest(text string) {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(text, time.Now())
	spew.Dump(err)
	if err == nil {
		spew.Dump(r.Time)
	}
}

func parse(args []string) string {
	var theTime string
	app.Version(version)

	hownow := New(time.Now(), *offset)

	switch kingpin.MustParse(app.Parse(args)) {
	case free.FullCommand():
		whenTest(*freeText)
	case bod.FullCommand():
		theTime = hownow.Process("bod", *epoch)
	case eod.FullCommand():
		theTime = hownow.Process("eod", *epoch)
	case bow.FullCommand():
		theTime = hownow.Process("bow", *epoch)
	case eow.FullCommand():
		theTime = hownow.Process("eow", *epoch)
	case bom.FullCommand():
		theTime = hownow.Process("bom", *epoch)
	case eom.FullCommand():
		theTime = hownow.Process("eom", *epoch)
	}

	return format(theTime, !*noNewLine)
}

func format(in string, newLine bool) string {
	if newLine {
		return fmt.Sprintf("%s\n", in)
	} else {
		return in
	}
}
