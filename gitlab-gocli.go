package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	parseConfig()
	app := cli.NewApp()
	app.Name = "gitlab-gocli"
	app.Version = Version
	app.Usage = ""
	app.Author = "Maxime Pelletier"
	app.Email = "pelletiermaxime@gmail.com"
	app.Commands = Commands
	app.Flags = GlobalFlags

	app.Run(os.Args)
}
