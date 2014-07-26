package main

import (
	log "github.com/cihub/seelog"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	parseConfig()
	loadLogger()
	defer log.Flush()

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
