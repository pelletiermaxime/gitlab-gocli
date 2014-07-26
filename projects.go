package main

import (
	"encoding/json"
	// log "github.com/cihub/seelog"
	"github.com/codegangsta/cli"
)

type project struct {
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
	w := getTabWriter()

	var JSONProjects []byte
	if c.Bool("o") || c.Bool("owned") {
		JSONProjects = NewRequest("projects/owned")
	} else {
		JSONProjects = NewRequest("projects")
	}

	var projects []project
	json.Unmarshal(JSONProjects, &projects)
	printHeader(w, "ID", "Path with namespace", "Web URL")
	for _, project := range projects {
		printColumn(w, project.ID, project.PathWithNamespace, project.WebURL)
	}
	w.Flush()
}
