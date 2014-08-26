package main

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/codegangsta/cli"
	"net/url"
)

type projectInfo struct {
	ID              int         `json:"id"`
	Description     interface{} `json:"description"`
	DefaultBranch   string      `json:"default_branch"`
	Public          bool        `json:"public"`
	VisibilityLevel int         `json:"visibility_level"`
	SSHUrlToRepo    string      `json:"ssh_url_to_repo"`
	HTTPUrlToRepo   string      `json:"http_url_to_repo"`
	WebURL          string      `json:"web_url"`
	Owner           struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
	} `json:"owner"`
	Name                 string `json:"name"`
	NameWithNamespace    string `json:"name_with_namespace"`
	Path                 string `json:"path"`
	PathWithNamespace    string `json:"path_with_namespace"`
	IssuesEnabled        bool   `json:"issues_enabled"`
	MergeRequestsEnabled bool   `json:"merge_requests_enabled"`
	WikiEnabled          bool   `json:"wiki_enabled"`
	SnippetsEnabled      bool   `json:"snippets_enabled"`
	CreatedAt            string `json:"created_at"`
	LastActivityAt       string `json:"last_activity_at"`
	Namespace            struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		OwnerID     int    `json:"owner_id"`
		Path        string `json:"path"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"namespace"`
	Permissions struct {
		ProjectAccess struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"project_access"`
		GroupAccess struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_access"`
	} `json:"permissions"`
	Archived bool `json:"archived"`
}

func doProjectInfo(c *cli.Context) {

	// projectID := url.QueryEscape(c.Args().First())
	projectID := c.Args().First()
	var JSONProjectInfo []byte
	JSONProjectInfo = NewRequest("projects/" + url.QueryEscape(projectID))
	// fmt.Println("projects/" + projectID)
	// fmt.Println("JSONProjectInfo:" + string(JSONProjectInfo))

	var project projectInfo
	err := json.Unmarshal(JSONProjectInfo, &project)
	if err != nil {
		log.Errorf("Error: %s for project ID %s with Body: %s", err, projectID, string(JSONProjectInfo))
	}
	// fmt.Printf("projects: %s", project)
	if project.ID == 0 {
		log.Criticalf("Project ID \"%s\" doesn't exist.", projectID)
		fmt.Println("")
		return
	}
	w := getTabWriter()
	printHeader(w, "ID", "Path with namespace", "Web URL", "Namespace ID")
	printLine(w, project.ID, project.PathWithNamespace, project.WebURL, project.Namespace.ID)
	w.Flush()
}

func doProjectCreate(c *cli.Context) {
	projectName := c.Args().First()
	v := url.Values{}
	v.Set("name", projectName)
	if c.Int("namespace-id") != 0 {
		v.Set("namespace_id", c.String("namespace-id"))
	}
	response := NewRequestPOST("projects", v)
	log.Debug(string(response))
}

func doProjectDelete(c *cli.Context) {
	projectID := c.Args().First()
	response := NewRequestDELETE("projects/" + projectID)
	log.Debug(string(response))
}
