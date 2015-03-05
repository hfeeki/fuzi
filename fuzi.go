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
	app.Email = "hfeeki@gmail.com"
	app.Action = doMain

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	app.Commands = Commands

	app.Run(os.Args)
}

func doMain(c *cli.Context) {
	name := "someone"
	if len(c.Args()) > 0 {
		name = c.Args()[0]
	}
	if c.String("lang") == "spanish" {
		println("Hola", name)
	} else if c.String("lang") == "english" {
		println("Hello", name)
	} else if c.String("lang") == "chinese" {
		println("Ni hao", name)
	}
}
