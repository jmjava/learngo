package composite

import "testing"

func TestComposite(t *testing.T){

	var branch = &Branch{Name: "branch 1"}

	var leaf1 = Leaflet{Name: "leaf 1"}
	var leaf2 = Leaflet{Name: "leaf 2"}

	//NOTE first way can pass without pointer
	//NOTE second way of creating branch with pointer
	//NOTE third way of creating branch with new (use pointer to pass)

	//var branch2 = Branch{Name: "branch 2"}  // can pass without pointer deref
	//var branch2 = &Branch{Name: "branch 2"} // need pointer deref
	var branch2 = new(Branch) // need pointer deref
	branch2.Name = "branch 2"

	branch.Add(leaf1)
	branch.Add(leaf2)
	//branch.AddBranch(branch2)
	branch.AddBranch(*branch2)

	branch.Perform()

}
