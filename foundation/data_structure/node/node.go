package node

import "fmt"

type Node struct {
	val  int
	next *Node
}

func insertToTail(head *Node, val int) (*Node, bool) {
	p := head
	if p == nil {
		return nil, false
	}

	for p != nil {
		if p.next == nil {
			tmpNode := &Node{
				val:  val,
				next: nil,
			}
			p.next = tmpNode
			return head, true
		}
		p = p.next
	}

	return nil, false
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

// 前提是，链表所有节点中的val唯一
func del(head *Node, val int) (*Node, bool) {
	p := head
	pre := p
	if p == nil {
		return nil, false
	}

	//head, middle, tail
	//head
	if head.val == val {
		//处理head指针未销毁问题
		head.val = head.next.val
		head.next = head.next.next
		return head, true
	}
	pre = p
	p = p.next
	//middle
	for p != nil {
		if p.val == val {
			pre.next = p.next
			return head, true
		}
		pre = pre.next
		p = p.next
	}
	//tail

	return nil, false
}

// return the new node pointer
func update(head *Node, val, target int) (*Node, bool) {
	p := head
	for p != nil {
		if p.val == val {
			p.val = target
			return head, true
		}
		p = p.next
	}

	return nil, false
}

func showNode(p *Node) {
	for p != nil {
		fmt.Printf("-> %v/%p\t", p.val, p)
		//fmt.Println(p)
		p = p.next
	}
	fmt.Println()
}

func find(p *Node, val int) (ret *Node, b bool) {

	for p != nil {
		if p.val == val {
			return p, true
		}
		p = p.next
	}

	return nil, false
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
