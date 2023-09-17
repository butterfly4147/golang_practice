package node_plus

import (
	"fmt"
	"testing"
)

/*
最终设计方案，参照golang的slice切片append设计方法
*/
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

func TestNode3(t *testing.T) {
	var head3 *Node
	println(head3 == nil)

	head3.append3(1)
	head3.append3(2)
	fmt.Println(head3)

	head3.traverse()
}

func TestNode4(t *testing.T) {
	var head3 *Node
	println(head3 == nil)

	head3.append3(1)
	head3.append3(2)
	fmt.Println(head3)

	head3.traverse()

	var sl []int
	sl = make([]int, 2)
	fmt.Printf("%v\n", sl)
	sl = append(sl, 1, 2)
	sl = append(sl, 1, 2, 3, 4, 5, 5, 2, 32, 2, 2, 2, 3)
	fmt.Printf("%v\n", sl)
}
