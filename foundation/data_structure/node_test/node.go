package main

import "fmt"

// 单向链表 头指针、头节点（方便链表增删的）、首元节点（第一个有效数据节点）
type Node struct {
	data int   //数据域
	next *Node //指针域
}

func ShowNode(p *Node) { //遍历链表值
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

// 查找对应值的index
func Find(p *Node, obj int) int { //获取目标值obj的index(0是头节点，找到的第一个值)，没找到返回-1
	head := p
	if head == nil {
		return -1
	}
	for i := 0; head != nil; i++ { //head != nil;  操作数据域
		if head.data == obj {
			return i
		}
		head = head.next
	}
	return -1
}

// 更新数据
func Update(p *Node, val, index int) bool { //修改目标index对应节点的值
	head := p
	if head == nil {
		return false
	}
	for i := 1; i <= index; i++ {
		head = head.next //略过头节点
		if head == nil {
			return false
		}
	}
	head.data = val

	return true
}

// 获取index对应的节点数据   获取失败返回-1
func GetNodeData(p *Node, index int) int {
	head := p
	if head == nil || index < 0 {
		return -1
	}
	for i := 1; i <= index; i++ {
		head = head.next //略过头节点
		if head == nil {
			return -1
		}
	}
	return head.data
}

// 获取前驱节点
func GetPreNode(p *Node, index int) *Node {
	head := p
	if head == nil || index < 0 {
		return head
	}
	//找到[前驱节点] 当前链表index位置上的前一个节点
	for ; index-1 > 0; index-- {
		if head.next != nil {
			head = head.next
		} else { //没找到
			return nil
		}
	}

	return head
}

// 插入节点 (0是头节点)
func Insert(p *Node, val, index int) bool {
	head := p
	if head == nil || index < 0 {
		return false
	}
	preNode := GetPreNode(head, index)
	if preNode == nil {
		return false
	}
	node := Node{data: val, next: preNode.next} //即使preNode.next=nil也无所谓
	preNode.next = &node

	return true
}

// 删除节点 (0是头节点)
func Delete(p *Node, val, index int) bool {
	head := p
	if head == nil || index < 0 {
		return false
	}
	preNode := GetPreNode(head, index)
	if preNode == nil {
		return false
	}
	preNode.next = preNode.next.next //不用判断 head.next.next==nil

	return true
}

func linkList(length int) *Node { // 1.尾插法 插入节点 head>>node1>>node2  <<tail
	//头节点
	head := Node{data: 0}
	var tail *Node
	tail = &head
	//构造单向链表
	for i := 1; i <= length; i++ {
		node := Node{data: i}
		tail.next = &node
		tail = &node
	}

	//返回头节点地址
	return &head
}

const LENGTH = 3

func main() {
	head := linkList(LENGTH)
	//遍历
	ShowNode(head)

	//获取某个值的索引
	fmt.Println("Find0:", Find(head, 0))
	fmt.Println("Find2:", Find(head, 2))
	fmt.Println("Find3:", Find(head, 4))
	//节点插入
	if Insert(head, 10, 4) {
		fmt.Println("插入数据")
		ShowNode(head)
	} else {
		fmt.Println("插入失败")
	}
	//删除节点
	if Delete(head, 10, 4) {
		fmt.Println("删除数据")
		ShowNode(head)
	} else {
		fmt.Println("删除失败")
	}
	//修改节点
	if Update(head, 22, 2) {
		fmt.Println("修改数据")
		ShowNode(head)
	} else {
		fmt.Println("修改失败")
	}
	//获取数据
	if data := GetNodeData(head, 2); data != -1 {
		fmt.Println("获取数据:", data)
	} else {
		fmt.Println("获取数据失败")
	}
	//头插法
	tail := linkList2(10)
	ShowNode(tail)
}

func linkList2(length int) *Node { // 2.头插法 插入节点  tail>> node2<<node1<<head
	head := Node{data: 0}
	var tail *Node
	tail = &head //tail用于记录头节点地址，首先指向头节点
	for i := 1; i < length; i++ {
		var node = Node{data: i}
		node.next = tail //将新创建的node的next指向头节点
		tail = &node     //重新赋值头节点
	}
	return tail
}
