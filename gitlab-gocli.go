package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitlab-gocli"
	app.Version = Version
	app.Usage = ""
	app.Author = "Maxime Pelletier"
	app.Email = "pelletiermaxime@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
