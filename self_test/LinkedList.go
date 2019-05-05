package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
   next *Node
   property int

}

func (linkedList *LinkedList) Length() int{
  length := 0
  for node := linkedList.head; node != nil ; node = node.next{
  	  length = length + 1
  }
  return length
}


func (linkedList *LinkedList) AddToHead(value int){
	node := Node{next: nil, property: value}

	node.next = linkedList.head

	linkedList.head = &node

}

func (linkedList *LinkedList) AddAfter(checkval int, value int) {
	node := Node{next: nil, property: value}
	nodeWith := linkedList.NodeWithValue(checkval)
	originalNext := nodeWith.next
	nodeWith.next = &node
	node.next = originalNext

}

func (linkedList *LinkedList) NodeWithValue (property int) *Node {
	var node *Node //for the loop
	var nodeWith *Node  // to return
	for node = linkedList.head ; node != nil ; node = node.next{
		if node.property == property{
			nodeWith = node
			break;
		}
	}
	return nodeWith

}

func (linkedList *LinkedList) AddToTail(value int) {

	tail := linkedList.Tail()
	tail.next = &Node{next : nil, property: value}

}

func (linkedList *LinkedList) Head(property int) *Node {

	return linkedList.head
}

func (linkedList *LinkedList) Tail() *Node {
    var node *Node
	for node = linkedList.head ; node != nil ; node = node.next{
		if node.next == nil {
			break
		}
	}
	return node

}

func (linkedList *LinkedList)Iterate () {
	fmt.Println("List")
	var node *Node
	for node = linkedList.head ; node != nil ; node = node.next{
		fmt.Println("Node Valule: ", node.property)
	}


}

func (linkedList *LinkedList ) PopFromHead()  *Node{
	var nodeToReturn *Node
	if (linkedList.head == nil){
		return nil
	}
	if (linkedList.head.next == nil){
		nodeToReturn = linkedList.head
		linkedList.head = nil
    }

	nodeToReturn = linkedList.head
	linkedList.head = nodeToReturn.next

	return nodeToReturn
}

func (linkedList *LinkedList ) PopFromTail()  *Node{
	tail := linkedList.Tail()
    var node *Node
	iter := 1
	//peak ahead to pop from the tail
	for node = linkedList.head; iter < linkedList.Length(); node = node.next{
		//peak ahead
		next := node.next
		fmt.Println(next)

		if (node.next.property == tail.property){
			node.next = nil
		}
		iter = iter + 1
	}

    return tail
}




func main(){

	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToTail(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	linkedList.Iterate()
	linkedList.PopFromHead()
	linkedList.Iterate()
	linkedList.PopFromTail()
	linkedList.Iterate()
	linkedList.PopFromTail()
	linkedList.Iterate()




}