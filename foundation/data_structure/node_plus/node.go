package node_plus

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	val  interface{}
	next *Node
}

// 需要方法接收器n!=nil
func (n *Node) append(val interface{}) error {
	head := n

	if head == nil {
		log.Printf("\tFirstly, please init the node: %v\n", head)
		return nil
	}

	for head != nil {
		if head.next == nil {
			head.next = &Node{
				val:  val,
				next: nil,
			}
			return nil
		}
		head = head.next
	}

	return errors.New("append failed! ")
}

// 类似与slice中的append使用
func (n *Node) append2(val interface{}) (*Node, error) {
	head := n

	if head == nil {
		//todo: if you want to use the form, you must to get a way to change the value of "var head *Node" which is as a method receiver
		//	which is mean that change the value pointed by the pointer
		head = &Node{
			val:  val,
			next: nil,
		}
		//log.Printf("\tFirstly, please init the node: %v\n", head)

		return head, nil
	}

	for head != nil {
		if head.next == nil {
			head.next = &Node{
				val:  val,
				next: nil,
			}
			return head, nil
		}
		head = head.next
	}

	return head, errors.New("append failed! ")
}

func (n *Node) traverse() {
	head := n
	if head == nil {
		fmt.Println("node is nil. ")
	}

	for head != nil {
		fmt.Printf("%v -> \t", head)
		head = head.next
	}

	fmt.Println("\n\ttraverse complete. ")
}
