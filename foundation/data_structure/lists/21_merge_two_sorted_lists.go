package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode

func init() {
}

func add(node *ListNode, val int) bool {
	if head == nil {
		head = &ListNode{val, nil}
		return true
	}

	if node.Val == val {
		fmt.Printf("节点已经存在：%d\n", val)
		return false
	}

	if node.Next == nil {
		node.Next = &ListNode{val, nil}
		return true
	}

	return add(node.Next, val)
}

func del(node *ListNode) bool {

	return true
}

func update(node *ListNode) bool {

	return true
}

func sel(node *ListNode) (ret *ListNode) {

	return nil
}

// traverse node one by one
func traverse(node *ListNode) {
	if node == nil {
		fmt.Printf("空链表。")
	}
	for node != nil {
		fmt.Printf("%d -> \t", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	return nil
}

func main() {
	//先把目前实现链表的方式背下来
	add(head, 1)
	add(head, 3)
	add(head, 2)
	traverse(head)

	mergeTwoLists(nil, nil)
}
