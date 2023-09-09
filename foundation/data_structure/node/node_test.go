package node

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	head := linkList(3)
	//traverse
	showNode(head)
	fmt.Println(find(head, 2))
	_, okInsert := insertToTail(head, 5)
	if okInsert {
		showNode(head)
	}
	//head是旧变量只是重新赋值；okDel是新变量
	_, okDel := del(head, 0)
	if okDel {
		showNode(head)
	}
	_, okUpdate := update(head, 2, 6)
	if okUpdate {
		showNode(head)
	}

	println("---------------")
	nodes := linkedNode([]int{2, 34, 4})
	showNode(nodes)

	println("---------------")
	nodes2 := linkedNode2([]int{2, 34, 4})
	showNode(nodes2)

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
