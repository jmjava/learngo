package decorator

import "testing"

func TestDecorator(t *testing.T){

	var process = &ProcessClass{}

	var decorator = &ProcessDecorator{}

	decorator.Process()

	decorator.ProcessInstance = process

	decorator.Process()

}
