package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/mattn/go-gtk/gtk"
	"strconv"
)

func doGUI(c *cli.Context) {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Gitlab GUI")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	scrolledwin := gtk.NewScrolledWindow(nil, nil)
	textview := gtk.NewTextView()
	textview.SetEditable(false)
	textview.SetCursorVisible(false)
	scrolledwin.Add(textview)
	vbox.Add(scrolledwin)

	buffer := textview.GetBuffer()

	button := gtk.NewButtonWithLabel("List projects")
	button.SetTooltipMarkup("List projects you have access to.")
	button.Clicked(func() {
		var iter gtk.TextIter
		buffer.GetStartIter(&iter)

		var JSONProjects []byte
		JSONProjects = NewRequest("projects")

		var projects []project
		json.Unmarshal(JSONProjects, &projects)
		var projectText string
		for _, project := range projects {
			projectText += strconv.Itoa(project.ID) + "-" + project.PathWithNamespace + " " + project.WebURL + "\n"
		}
		buffer.Insert(&iter, projectText)
	})

	vbox.PackEnd(button, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(800, 500)
	window.ShowAll()
	gtk.Main()
}
