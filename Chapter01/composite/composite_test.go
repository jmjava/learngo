package composite

import "testing"

func TestComposite(t *testing.T){

	var branch = &Branch{Name: "branch 1"}

	var leaf1 = Leaflet{Name: "leaf 1"}
	var leaf2 = Leaflet{Name: "leaf 2"}

	var branch2 = Branch{Name: "branch 2"}

	branch.Add(leaf1)
	branch.Add(leaf2)
	branch.AddBranch(branch2)

	branch.Perform()


}
