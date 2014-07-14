package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"os"
	"text/tabwriter"
)

type Project struct {
	Archived             bool   `json:"archived"`
	CreatedAt            string `json:"created_at"`
	DefaultBranch        string `json:"default_branch"`
	Description          string `json:"description"`
	HTTPURLToRepo        string `json:"http_url_to_repo"`
	ID                   int    `json:"id"`
	IssuesEnabled        bool   `json:"issues_enabled"`
	LastActivityAt       string `json:"last_activity_at"`
	MergeRequestsEnabled bool   `json:"merge_requests_enabled"`
	Name                 string `json:"name"`
	NameWithNamespace    string `json:"name_with_namespace"`
	Namespace            struct {
		Avatar struct {
			URL interface{} `json:"url"`
		} `json:"avatar"`
		CreatedAt   string      `json:"created_at"`
		Description string      `json:"description"`
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		OwnerID     interface{} `json:"owner_id"`
		Path        string      `json:"path"`
		UpdatedAt   string      `json:"updated_at"`
	} `json:"namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
	Public            bool   `json:"public"`
	SnippetsEnabled   bool   `json:"snippets_enabled"`
	SSHURLToRepo      string `json:"ssh_url_to_repo"`
	VisibilityLevel   int    `json:"visibility_level"`
	WebURL            string `json:"web_url"`
	WikiEnabled       bool   `json:"wiki_enabled"`
}

func doProjects(c *cli.Context) {
	info := color.New(color.FgYellow, color.BgBlack).SprintFunc()
	infoBold := color.New(color.FgYellow, color.BgBlack, color.Bold).SprintFunc()
	normal := color.New(color.FgWhite, color.BgBlack).SprintFunc()
	normalBold := color.New(color.FgWhite, color.BgBlack, color.Bold).SprintFunc()
	w := new(tabwriter.Writer)
	// Format in space-separated columns of minimal width 8
	w.Init(os.Stdout, 8, 0, 3, ' ', 0)

	var JSONProjects []byte
	if c.Bool("o") || c.Bool("owned") {
		JSONProjects = NewRequest("projects/owned")
	} else {
		JSONProjects = NewRequest("projects")
	}

	var projects []Project
	json.Unmarshal(JSONProjects, &projects)
	fmt.Fprintln(w, infoBold("ID"), "\t", info("Path with namespace"), "\t", info("Web URL"))
	// fmt.Println(projects)
	for _, project := range projects {
		fmt.Fprintln(w, normalBold(fmt.Sprintf("%v", project.ID)), "\t", normal(project.PathWithNamespace), "\t", normal(project.WebURL))
	}
	w.Flush()
}
