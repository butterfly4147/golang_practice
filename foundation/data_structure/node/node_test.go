package node

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	head := linkList(3)
	//traverse
	ShowNode(head)

	println("---------------")
	nodes := linkedNode([]int{2, 34, 4})
	ShowNode(nodes)
	println("---------------")
	nodes2 := linkedNode2([]int{2, 34, 4})
	ShowNode(nodes2)
}

func TestPointer(t *testing.T) {
	head := Node{
		val:  1,
		next: nil,
	}
	fmt.Println("tail: ", head)
	var tail *Node
	fmt.Println("tail: ", tail)
	tail = &head
	fmt.Println(tail)
}
