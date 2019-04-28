package main

import "learngo/Chapter01/adapter"
import "learngo/Chapter01/bridge"
import "learngo/Chapter01/composite"
import "learngo/Chapter01/decorator"

// main method
func main() {

	runAdapter()
	runBridge()
	runComposite()
	runDecorator()

}

func runAdapter(){

	//note have to make method public

	var processor adapter.IProcess = adapter.Adapter{}

	processor.Process()
}

func runBridge() {

	    //note have to make fields and methods public to be able to assign from this package

		var x = [5]float32{1, 2, 3, 4, 5}
		var y = [5]float32{1, 2, 3, 4, 5}
		var contour bridge.IContour = bridge.DrawContour{x, y, bridge.DrawShape{}, 2}

		contour.DrawContour(x, y)
		contour.ResizeByFactor(2)


}

func runComposite() {

	var branch = &composite.Branch{Name: "branch 1"}

	var leaf1 = composite.Leaflet{Name: "leaf 1"}
	var leaf2 = composite.Leaflet{Name: "leaf 2"}

	var branch2 = composite.Branch{Name: "branch 2"}

	branch.Add(leaf1)
	branch.Add(leaf2)
	branch.AddBranch(branch2)

	branch.Perform()

}

func runDecorator() {

	var process = &decorator.ProcessClass{}

	var decorator = &decorator.ProcessDecorator{}

	decorator.Process()

	decorator.ProcessInstance = process

	decorator.Process()


}
