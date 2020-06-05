package datastruct

import "fmt"

type SingleCircleLink struct {
	// linklist.go中定义的节点
	head *Node
}

func NewSingleCircleLink(num int) *SingleCircleLink {
	if num < 0 {
		panic("num必须大于等于0")
		return nil
	}
	head := &Node{name: "**head**"}
	if num == 0 {
		// 空链表，将头结点的next指向自己.循环链表中，head结点是不算在节点内的
		head.next = head
	} else {
		tmp := head
		for i := 1; i <= num; i++ {
			nodeName := fmt.Sprintf("circleNode-%d", i)
			newNode := Node{no: i, name: nodeName}
			tmp.next = &newNode
			newNode.next = head.next
			tmp = &newNode
		}
	}
	return &SingleCircleLink{head: head}
}

func (sc *SingleCircleLink) Length() int {
	head := sc.head
	if head.next == head {
		return 0
	}
	// 移动获取第一个结点
	head = head.next
	count := 1
	for {
		if head.next == sc.head.next {
			break
		}
		count++
		head = head.next
	}
	return count
}

func (sc *SingleCircleLink) Show() {
	head := sc.head.next
	if sc.head.next == sc.head {
		fmt.Println("Empty Single Circle Link")
		return
	}
	for {
		fmt.Println("node name ", head.name, "node.next name ", head.next.name)
		if head.next == sc.head.next {
			break
		}
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
	// 先判断位置是否合法。
	if n < 0 || n > sc.Length() + 1 {
		fmt.Println("位置不合法")
		return
	}
	head := sc.head
	// 说明这是第一个结点，因为第一个结点的next是自己，特殊处理之
	if n == 1 && sc.Length() == 0 {
		head.next = &newNode
		newNode.next = &newNode
	} else {
		// 插入位置为第一个，但链表至少存在一个结点
		if n == 1 {
			// 将原来链表的下一个指向为新结点，因为是循环链表
			head.next.next = &newNode
			newNode.next = head.next
			head.next = &newNode
		} else {
			// 常规插入，先获取待插入节点的前一个节点
			preNode := sc.Get(n - 1)
			newNode.next = preNode.next
			preNode.next = &newNode
		}
	}

}

func (sc *SingleCircleLink) Delete(n int) {
	// 位置合法判断
	if n <= 0 || n > sc.Length() {
		fmt.Println("位置不合法")
		return
	}
	// 如果删除的是第一个节点，需要特殊处理。 获取待删除前一个结点为最后的结点
	// 修改head.next 与最后一个节点的next
	if n == 1 {
		preNode := sc.Get(sc.Length())
		deleteNode := sc.Get(1)
		preNode.next = deleteNode.next
		sc.head.next = deleteNode.next
	} else {
		// 获取待删除的前一个结点
		preNode := sc.Get(n - 1)
		// 获取待删除节点
		deleteNode := sc.Get(n)
		// 将待删除前节点的下一个结点指向待删除的下一个结点
		preNode.next = deleteNode.next
	}
}

// 模拟约瑟夫环出圈
func (sc *SingleCircleLink) Out(start, count, total int)  int {
	// 先对输入数据做校验
	if start < 1 || count > total || sc.Length() != total || start > total {
		fmt.Println("输入不合法")
		return 0
	}
	// 先定位到start的元素，以及获取它前一个节点，作为辅助指针
	var helper *Node
	if start == 1 {
		helper = sc.Get(sc.Length())
	} else {
		helper = sc.Get(start - 1)
	}
	curr := sc.Get(start)
	// helper与curr指针同时移动count - 1次（因为它自身也算一次），并将移动后的节点出列
	for {
		// 如果前一个与后一个相等，说明队列最后只有一个节点了
		if curr == helper {
			break
		}
		cnt := count - 1
		// 内部的循环中，让helper与curr指针移动count - 1次，因为它本身也算一次
		for {
			if cnt == 0 {
				// 达到指定次数后，退出内层循环
				break
			}
			helper = helper.next
			curr = curr.next
			cnt --
		}
		// 内循环退出后，表示curr指针已经移动到了待删除的节点。然后删除该节点，将它上层的next指向自己的next
		helper.next = curr.next
		fmt.Println("出圈号为 ", curr.no)
		curr = curr.next
	}
	return curr.no
}