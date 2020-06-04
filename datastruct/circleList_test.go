package datastruct

import (
	"fmt"
	"testing"
)

func TestNewSingleCircleLink(t *testing.T) {
	singleCircleLink := NewSingleCircleLink()
	singleCircleLink.Show()
	length := singleCircleLink.Length()
	fmt.Println(length)
	node := Node{name: "new node"}
	singleCircleLink.Insert(1, node)
	node.name = "another name"
	singleCircleLink.Insert(1, node)
	fmt.Println("after insert show")
	singleCircleLink.Show()
	fmt.Println(singleCircleLink.Length())
	fmt.Println("after delete node ")
	singleCircleLink.Delete(1)
	singleCircleLink.Show()
	fmt.Println(singleCircleLink.Length())
}
