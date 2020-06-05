package datastruct

import (
	"fmt"
	"testing"
)

func TestNewSingleCircleLink(t *testing.T) {
	fmt.Println("创建一个空的循环链表")
	singleCircleLink := NewSingleCircleLink(0)
	singleCircleLink.Show()
	fmt.Println("当前循环链表长度为", singleCircleLink.Length(), "往循环链表中插入一个节点")
	node := Node{no: 1, name: "insert node"}
	singleCircleLink.Insert(1, node)
	fmt.Println("插入一个后的循环链表长度为", singleCircleLink.Length())
	node.name = "another node"
	singleCircleLink.Insert(1, node)
	fmt.Println("插入一个后的循环链表长度为", singleCircleLink.Length())
	node.name = "3th node"
	singleCircleLink.Insert(3, node)
	singleCircleLink.Show()
	singleCircleLink.Delete(3)
	fmt.Println("删除一个结点后的情况", singleCircleLink.Length())
	singleCircleLink.Show()

}

// 使用链表模拟约瑟夫环问题
func TestSingleCircleLink_Out(t *testing.T) {
	circleLink := NewSingleCircleLink(5)
	circleLink.Show()
	fmt.Println(circleLink.Length())
	fmt.Println("约瑟夫环出圈...")
	last := circleLink.Out(1, 2, 5)
	fmt.Println("最后出圈号码为 ", last)
}
