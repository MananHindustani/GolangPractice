package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("chat", "A command-line chat application.")
)

func main() {
	app.HelpFlag.Short('h')
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	}
}
