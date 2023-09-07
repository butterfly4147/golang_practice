package node

import "fmt"

type Node struct {
	val  int64
	next *Node
}

func ShowNodeList(p *Node) {
	if p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

func linkList(len int) *Node {
	head := Node{0, nil}
	//todo：tail := &head
	var tail *Node
	tail = &head
	for i := 0; i < len; i++ {

	}

	return &head
}
