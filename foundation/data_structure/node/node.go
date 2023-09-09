package node

import "fmt"

type Node struct {
	val  int
	next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Printf("-> %v/%p\t", p.val, p)
		//fmt.Println(p)
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
		node := Node{val: i}

		tail.next = &node
		tail = &node
	}
	//return the head pointer
	return head.next
}

func linkedNode(nums []int) (ret *Node) {
	//init
	ret = &Node{
		val:  0,
		next: nil,
	}
	//assign a pointer value to tail
	tail := ret
	for _, num := range nums {
		node := &Node{
			val:  num,
			next: nil,
		}

		tail.next = node
		tail = node
	}

	return ret.next
}

// 减少初始化空节点。不是钻牛角尖，只是尝试另一种实现方式
func linkedNode2(nums []int) (ret *Node) {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	//init
	ret = &Node{
		val:  nums[0],
		next: nil,
	}
	//assign a pointer value to tail
	tail := ret
	for idx, num := range nums {
		if idx == 0 {
			continue
		}

		node := &Node{
			val:  num,
			next: nil,
		}

		tail.next = node
		tail = node
	}

	return ret
}

//func linkList(length int) *Node { // 1.尾插法 插入节点 head>>node1>>node2  <<tail
//	//头节点
//	head := Node{val: 0}
//	var tail *Node
//	tail = &head
//	//构造单向链表
//	for i := 1; i <= length; i++ {
//		node := Node{val: i}
//		tail.next = &node
//		tail = &node
//	}
//
//	//返回头节点地址
//	return &head
//}
