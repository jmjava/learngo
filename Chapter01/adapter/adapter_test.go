package adapter

import "testing"

func TestAdapter(t *testing.T){

	processor := IProcess(Adapter{})

	processor.Process()

}