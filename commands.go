package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandToken,
	commandExpire,
	commandRegister,
	commandProfile,
	commandChgpwd,
}

var commandToken = cli.Command{
	Name:  "token",
	Usage: "",
	Description: `
`,
	Action: doToken,
}

var commandExpire = cli.Command{
	Name:  "expire",
	Usage: "",
	Description: `
`,
	Action: doExpire,
}

var commandRegister = cli.Command{
	Name:  "register",
	Usage: "",
	Description: `
`,
	Action: doRegister,
}

var commandProfile = cli.Command{
	Name:  "profile",
	Usage: "",
	Description: `
`,
	Action: doProfile,
}

var commandChgpwd = cli.Command{
	Name:  "chgpwd",
	Usage: "",
	Description: `
`,
	Action: doChgpwd,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doToken(c *cli.Context) {
}

func doExpire(c *cli.Context) {
}

func doRegister(c *cli.Context) {
}

func doProfile(c *cli.Context) {
}

func doChgpwd(c *cli.Context) {
}
