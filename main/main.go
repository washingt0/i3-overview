package main

import (
	"log"

	i3 "github.com/washingt0/i3-overview"
)

func main() {
	tree := i3.GetTree()

	var apps []i3.Application
	for _, w := range i3.GetOnlyWorkspaces(tree) {
		i3.GetAllApplications(&w, &apps, 0)
	}

	for i := 0; i < len(apps); i++ {
		log.Printf("%+v\n\n", apps[i])
	}

}
