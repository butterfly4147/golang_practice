package node_plus

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	head := &Node{
		val:  "head",
		next: nil,
	}

	_ = head.append(1)
	_ = head.append(2)

	head.traverse()
}

func TestNode2(t *testing.T) {
	var head2 *Node
	println(head2 == nil)

	head2, _ = head2.append2(1)
	head2, _ = head2.append2(2)
	fmt.Println(head2)

	head2.traverse()
}
