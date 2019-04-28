//main package has examples shown
// in Go Data Structures and algorithms book
package composite

// importing fmt package
import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	Perform()
}

// Leaflet struct
type Leaflet struct {
	Name string
}

// Leaflet class method Perform
func (leaf *Leaflet) perform() {

	fmt.Println("Leaflet " + leaf.Name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	Name     string
	branches []Branch

}

// Branch class method Perform
func (branch *Branch) Perform() {

	fmt.Println("Branch: " + branch.Name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.Perform()
	}
}

// Branch class method Add leaflet
func (branch *Branch) Add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)

}

//Branch class method AddBranch branch
func (branch *Branch) AddBranch(newBranch Branch) {

	branch.branches = append(branch.branches, newBranch)
}

//Branch class  method getLeaflets
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// main method
func main() {


}
