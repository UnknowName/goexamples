package datastruct

import (
	"fmt"
	"testing"
)

func TestNewDLink(t *testing.T) {
	dlink := NewDLink(3)
	dlink.Show()
	lenght := dlink.Length()
	fmt.Println(lenght)
	newNode := DNode{name: "insert node"}
	dlink.Insert(2, newNode)
	dlink.Show()
	fmt.Println(dlink.Length())
}