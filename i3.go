package i3overview

import (
	"encoding/json"
	"log"
	"os/exec"
)

type windowProperties struct {
	Class    string `json:"class"`
	Title    string `json:"title"`
	Instance string `json:"instance"`
}

type i3Node struct {
	ID               int64            `json:"id"`
	Type             string           `json:"type"`
	Focused          bool             `json:"focused"`
	Layout           string           `json:"layout"`
	Name             string           `json:"name"`
	Nodes            []i3Node         `json:"nodes"`
	WindowProperties windowProperties `json:"window_properties"`
	Output           string           `json:"output"`
	Number           int              `json:"num"`
}

func GetTree() *i3Node {
	dump, err := exec.Command("i3-msg", "-t", "get_tree").Output()
	if err != nil {
		log.Fatal(err)
	}

	var tree i3Node

	err = json.Unmarshal(dump, &tree)
	if err != nil {
		log.Fatal(err)
	}

	return &tree
}

func GetOnlyWorkspaces(tree *i3Node) []i3Node {
	var workspaces []i3Node
	for _, n := range tree.Nodes {
		if n.Type == "output" && n.Name == "__i3" {
			continue
		}
		for _, m := range n.Nodes {
			if m.Type == "con" {
				for _, o := range m.Nodes {
					if o.Type == "workspace" {
						workspaces = append(workspaces, o)
					}
				}
			}
		}
	}
	return workspaces
}
