package datastruct

import "fmt"

type SingleCircleLink struct {
	// linklist.go中定义的节点
	head *Node
}

func NewSingleCircleLink() *SingleCircleLink {
	head := &Node{name: "**head**"}
	// 空链表，将头结点的next指向自己
	head.next = head
	return &SingleCircleLink{head: head}
}

func (sc *SingleCircleLink) Length() int {
	head := sc.head
	count := 0
	for {
		if head.next == sc.head {
			break
		}
		count++
		head = head.next
	}
	return count
}

func (sc *SingleCircleLink) Show() {
	count := 0
	head := sc.head
	if sc.head.next == sc.head {
		fmt.Println("Empty Single Circle Link")
		return
	}
	for {
		fmt.Println("node name ", head.name)
		if head.next == sc.head {
			break
		}
		count++
		head = head.next
	}
}

func (sc *SingleCircleLink) Get(n int) *Node {
	if n <= 0 || n > sc.Length() {
		fmt.Println("位置不合法")
		return nil
	}
	head := sc.head
	count := 0
	for {
		if count == n {
			break
		}
		head = head.next
		count++
	}
	return head
}

func (sc *SingleCircleLink) Insert(n int, newNode Node) {
	// 为1时，表示在头结点插入
	head := sc.head
	if n == 1 {
		newNode.next = head.next
		head.next = &newNode
		return
	}
	// 因为是单链表，需要拿到它的前一个
	preNode := sc.Get(n - 1 )
	newNode.next = preNode.next
	preNode.next = &newNode
}

func (sc *SingleCircleLink) Delete(n int) {
	// 位置合法判断
	if n <= 0 || n > sc.Length() {
		fmt.Println("位置不合法")
		return
	}
	// 如果删除的是第一个节点，需要特殊处理
	if n == 1 {
		head := sc.head
		head.next = head.next.next
		return
	}
	// 获取待删除的前一个结点
	preNode := sc.Get(n - 1)
	// 获取待删除节点
	deleteNode := sc.Get(n)
	// 将待删除前节点的下一个结点指向待删除的下一个结点
	preNode.next = deleteNode.next
}