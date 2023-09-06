package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: 空闲了可以试试构造链表节点的CRUD操作
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else { //include ">="
			cur.Next = list2
			list2 = list2.Next
		}
		//这一步是为了将cur指向尾节点。因为上述if向cur指向的尾节点后新增了一个节点，形成了cur不是尾节点的情况，需要这一步来进行调整
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return dump.Next
}

func main() {
	mergeTwoLists(nil, nil)
}
