package structs

import (
    "fmt"
)

type LRUCache struct {
    dict map[int]*DNode
    max int
    cache *DLink
}

type DLink struct {
    Head *DNode
    Tail *DNode
}

func (dl *DLink) Show() {
    tmp := dl.Head.next
    for tmp.next != nil {
        // fmt.Printf("k=%d v=%d %p\n", tmp.key, tmp.value, tmp)
        tmp = tmp.next
    }
}

func (dl *DLink) AddNodeToTail(node *DNode) {
    oldPre := dl.Tail.pre
    oldPre.next = node
    node.next = dl.Tail
    node.pre = oldPre
    dl.Tail.pre = node
}

func (dl *DLink) Delete() *DNode {
    // 删除最前面的节点
    dl.Head = dl.Head.next
    return dl.Head
}

func (dl *DLink) MoveToTail(node *DNode) {
    // 原来的前驱节点
    oldPre := node.pre
    // fmt.Println("pre ", oldPre, )
    // next 1 --> 2 --> 3 ===> 1 ---> 3
    oldPre.next = node.next
    // pre 1 <--- 2 <--- 3 ===> 1 <--- 3
    node.next.pre = oldPre
    // 将移出的结点追加进尾部
    dl.AddNodeToTail(node)
}


// 双链表的节点

type DNode struct {
    key int
    value int
    pre   *DNode
    next  *DNode
}



func Constructor(capacity int) LRUCache {
    l := LRUCache{
        dict: make(map[int]*DNode),
        max: capacity,
        cache: &DLink{Tail: new(DNode), Head: new(DNode)},
    }
    l.cache.Tail.pre = l.cache.Head
    l.cache.Head.next = l.cache.Tail
    return l
}

func (this *LRUCache) Show() {
    fmt.Println(this.dict, this.max)
    this.cache.Show()
}

func (this *LRUCache) Get(key int) int {
    node,  _ := this.dict[key]
    if node == nil {
        return -1
    }
    this.cache.MoveToTail(node)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    var node *DNode
    node, _ = this.dict[key]
    if node == nil {
        // 查下当前是否还有空间可以存放
        if this.max == 0 {
            deleteNode := this.cache.Delete()
            fmt.Println("deleteNode ", deleteNode.key)
            delete(this.dict, deleteNode.key)
            this.max += 1
        }
        node = new(DNode)
        node.value = value
        node.key = key
        this.max -= 1
        this.dict[key] = node
        this.cache.AddNodeToTail(node)
    } else {
        node.value = value
        this.cache.MoveToTail(node)
    }
}

