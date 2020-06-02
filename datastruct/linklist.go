package datastruct

import "fmt"

type Node struct {
	no   int
	name string
	next *Node
}

type LinkList struct {
	head *Node
}

func NewLinkList() *LinkList {
	head := &Node{}
	head.next = nil
	return &LinkList{head: head}
}

func (l *LinkList) Init(num int, reserve bool) {
	if reserve {
		// head是头结点，它不存储数据，并不是有效数据的一部分
		head := l.head
		for i := 1; i <= num; i++ {
			nodeName := fmt.Sprintf("name-%d", i)
			node := Node{no: i, name: nodeName}
			node.next = head.next
			head.next = &node
		}
	} else {
		// 使用一个辅助变量，标示尾结点
		tail := l.head
		for i := 1; i <= num; i++ {
			nodeName := fmt.Sprintf("name-%d", i)
			node := Node{no: i, name: nodeName}
			tail.next = &node
			tail = &node
		}
	}
}

func (l *LinkList) Show() {
	firstNode := l.head.next
	if firstNode == nil {
		fmt.Println("Empty LinkList")
		return
	}
	for {
		fmt.Println("Node no ", firstNode.no, " Node name ", firstNode.name)
		if firstNode.next == nil {
			break
		}
		firstNode = firstNode.next
	}
}

func (l *LinkList) Length() int {
	count := 0
	firstNode := l.head
	for {
		if firstNode.next == nil {
			break
		}
		count++
		firstNode = firstNode.next
	}
	return count
}

func (l *LinkList) Insert(i int, node Node) {
	if i <= 0 || i > l.Length() {
		panic("值不合法")
	}
	// 值为1时，插入在链表最顶部
	if i == 1 {
		node.next = l.head.next
		l.head.next = &node
	} else {
		// 正常插入，需要先通过头结点遍历到指定i节点的前一个节点的位置，再将新节点插入
		count := 1
		preNode := l.head.next
		for {
			if count + 1 == i {
				break
			}
			preNode = preNode.next
			count++
		}
		// 到达指定i节点的前一个节点，如果不在它的前一节点插入，会丢失后续节点的信息
		node.next = preNode.next
		preNode.next = &node
	}
}

func (l *LinkList) Append(node Node) {
	head := l.head
	for {
		if head.next == nil {
			break
		}
		head = head.next
	}
	head.next = &node
	node.next = nil
}

func (l *LinkList) Get(n int) *Node {
	count := 1
	node := l.head.next
	if n == 1 {
		return l.head.next
	}
	for {
		if count == n {
			break
		}
		node = node.next
		count++
	}
	return node
}

func (l *LinkList) AddByOrder(node Node) {
	head := l.head
	for {
		if head.next == nil {
			break
		}
		if head.next.no >= node.no {
			break
		}
		head = head.next
	}
	node.next = head.next
	head.next = &node
}

func (l *LinkList) Delete(n int) bool {
	if n <= 0 || n > l.Length() {
		panic("位置不合法")
	}
	// 先从头结点开始遍历至待删除结点的前一个结点
	head := l.head
	count := 1
	for {
		if count == n {
			break
		}
		head = head.next
		count++
	}
	//　跳出循环后，head指向的是待删除的前一个结点.因为待删除的节点就是head.next
	deleteNode := head.next
	// 将待删除前一个结点的下一个结点指向删除结点的下一个
	head.next = deleteNode.next
	// free(deleteNode)
	return true
}

// 将链表反转
func (l *LinkList) Reverse() {
	head := l.head.next
	reverseHead := Node{}
	reverseHead.next = nil
	for {
		if head == nil {
			break
		}
		tmp := head
		oldNext := head.next
		tmp.next = reverseHead.next
		reverseHead.next = tmp
		head = oldNext
	}
	l.head.next = reverseHead.next
}