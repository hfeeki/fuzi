package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fuzi"
	app.Version = Version
	app.Usage = ""
	app.Author = "Michael"
	app.Email = "hfeeki@outlook.com"
	app.Action = doMain
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
}
