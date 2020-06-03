package datastruct

import "fmt"

type DNode struct {
	pre  *DNode
	next *DNode
	name string
	no   int
}

type DLink struct {
	head *DNode
}

func NewDLink(n int) *DLink {
	head := &DNode{name: "head"}
	tmp := head
	if n <= 0 {
		return &DLink{head: head}
	} else {
		for i := 0; i < n; i++ {
			name := fmt.Sprintf("name-d%d", i)
			newNode := &DNode{no: i, name: name}
			tmp.next = newNode
			newNode.pre = tmp
			tmp = newNode
		}
	}
	return &DLink{head: head}
}

func (dl *DLink) Show() {
	head := dl.head
	node := head.next
	if node == nil {
		fmt.Println("Empty DLink")
	} else {
		for {
			if node == nil {
				break
			}
			fmt.Println("DLNode name ", node.name, " Before Node ", node.pre.name)
			node = node.next
		}
	}
}

func (dl *DLink) Delete(n int) {
	head := dl.head
	count := 0
	for {
		if head.next == nil || count == n{
			break
		}
		count++
		head = head.next
	}
	head.pre.next = head.next
	if head.next != nil {
		head.next.pre = head.pre
	}
}

func (dl *DLink) Length() uint32 {
	head := dl.head
	count := uint32(0)
	for {
		if head.next == nil {
			break
		}
		count++
		head = head.next
	}
	return count
}

func (dl *DLink) Get(n uint32) *DNode {
	if dl.Length() == uint32(0) {
		fmt.Println("Empty DLink")
		return nil
	}
	if n <= 0 || n > dl.Length() {
		fmt.Println("位置不合法")
		return nil
	}
	head := dl.head.next
	count := uint32(1)
	for {
		if head.next == nil || count == n {
			break
		}
		head = head.next
		count++
	}
	return head
}

func (dl *DLink) Insert(n uint32, newNode DNode) {
	node := dl.Get(n)
	// 新结点的前结点为老结点的前结点
	newNode.pre = node.pre
	// 老节点的前一个结点的下一个结点为新结点
	node.pre.next = &newNode
	// 新结点的下一个为老的
	newNode.next = node
	// 老的前一个结点为新的
	node.pre = &newNode
}