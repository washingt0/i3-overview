package i3overview

type Application struct {
	Workspace int
	Name      string
	Window    windowProperties
	Output    string
}

func GetAllApplications(tree *i3Node, output *[]Application, num int) {
	if tree.Number != 0 {
		num = tree.Number
	}

	if len(tree.Nodes) > 0 {
		for i := 0; i < len(tree.Nodes); i++ {
			GetAllApplications(&tree.Nodes[i], output, num)
		}
		return
	}

	*output = append(*output, Application{Workspace: num, Name: tree.Name, Window: tree.WindowProperties, Output: tree.Output})
}
