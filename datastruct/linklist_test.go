package datastruct

import (
	"fmt"
	"testing"
)

func TestNewLinkList(t *testing.T) {
	lst := NewLinkList()
	lst.Init(3, false)
	node := Node{no: 4, name: "append node"}
	lst.Append(node)
	nodes := lst.Get(1)
	nodes.name = "new name"
	lst.Show()
	lst.Reverse()
	fmt.Println("after Reverse")
	lst.Show()
}
