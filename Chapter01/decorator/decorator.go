//main package has examples shown
// in Go Data Structures and algorithms book
package decorator

// importing fmt package
import (
	"fmt"
)

// IProcess Interface
type IProcess interface {
	Process()
}

//ProcessClass struct
type ProcessClass struct{}

//ProcessClass method Process
func (process *ProcessClass) Process() {
	fmt.Println("ProcessClass Process")
}

//ProcessDecorator struct
type ProcessDecorator struct {
	ProcessInstance *ProcessClass
}

//ProcessDecorator class method Process
func (decorator *ProcessDecorator) Process() {
	if decorator.ProcessInstance == nil {
		fmt.Println("ProcessDecorator  Process")
	} else {
		fmt.Printf("ProcessDecorator  Process  and ")
		decorator.ProcessInstance.Process()

	}
}

//main method
func main() {

	var process = &ProcessClass{}

	var decorator = &ProcessDecorator{}

	decorator.Process()

	decorator.ProcessInstance = process

	decorator.Process()

}
