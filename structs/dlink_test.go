package structs

import (
    "fmt"
    "testing"
)

func TestDLink_AddNodeToTail(t *testing.T) {
    link := DLink{Tail: new(DNode), Head: new(DNode)}
    link.Tail.pre = link.Head
    link.Head.next = link.Tail
    secondNode := new(DNode)
    for i := 1; i <= 3; i++ {
        node := &DNode{key: i, value: i}
        link.AddNodeToTail(node)
        if i == 2 {
            secondNode = node
        }
    }
    fmt.Println("show")
    link.Show()
    fmt.Println("move node ", secondNode)
    link.MoveToTail(secondNode)
    link.Delete()
    link.Delete()
    link.Show()

}

func TestConstructor(t *testing.T) {
    cache := Constructor(2)
    cache.Put(2, 12)
    cache.Put(11,11)
    cache.Put(12,123)
    cache.Put(13,123)
    fmt.Println("get key", cache.Get(12))
    cache.Put(14,123)
    fmt.Println("get key", cache.Get(13))
    cache.Show()
    fmt.Println("ok------------------")
    cache.Put(13,123)
    cache.Show()
}
