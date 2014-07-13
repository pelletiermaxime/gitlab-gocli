package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandUsers,
	commandProjects,
	commandLogin,
	commandIssues,
}

var commandUsers = cli.Command{
	Name:  "users",
	Usage: "",
	Description: `
`,
	Action: doUsers,
}

var commandProjects = cli.Command{
	Name:  "projects",
	Usage: "",
	Description: `
`,
	Action: doProjects,
}

var commandLogin = cli.Command{
	Name:  "login",
	Usage: "",
	Description: `
`,
	Action: doLogin,
}

var commandIssues = cli.Command{
	Name:  "issues",
	Usage: "",
	Description: `
`,
	Action: doIssues,
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

func doUsers(c *cli.Context) {
}

func doProjects(c *cli.Context) {
}

func doLogin(c *cli.Context) {
}

func doIssues(c *cli.Context) {
}
