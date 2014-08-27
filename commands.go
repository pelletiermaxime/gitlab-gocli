package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	// commandUsers,
	commandProjects,
	// commandLogin,
	// commandIssues,
	commandProject,
	commandGui,
}

var GlobalFlags = []cli.Flag{}

var commandUsers = cli.Command{
	Name:  "users",
	Usage: "",
	Description: `
`,
	Action: doUsers,
}

var projectsFlags = []cli.Flag{
	cli.BoolFlag{Name: "owned, o", Usage: "Only projects owned by the authenticated user."},
}

var projectCreateFlags = []cli.Flag{
	cli.IntFlag{Name: "namespace-id, n", Usage: "Namespace ID."},
}

var commandProjects = cli.Command{
	Name:  "projects",
	Usage: "List projects",
	Description: `
`,
	Action: doProjects,
	Flags:  projectsFlags,
}

var commandProject = cli.Command{
	Name:  "project",
	Usage: "Create/delete/info a project",
	Description: `
`,
	Subcommands: []cli.Command{
		{
			Name:   "info",
			Usage:  "List all project info",
			Action: doProjectInfo,
		},
		{
			Name:   "create",
			Usage:  "Create a new project",
			Action: doProjectCreate,
			Flags:  projectCreateFlags,
		},
		{
			Name:   "delete",
			Usage:  "Delete a project",
			Action: doProjectDelete,
		},
	},
	// Action: doProject,
	// Flags:  ProjectFlags,
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

var commandGui = cli.Command{
	Name:  "gui",
	Usage: "Graphical interface",
	Description: `
`,
	Action: doGUI,
	// Flags:  projectsFlags,
}
