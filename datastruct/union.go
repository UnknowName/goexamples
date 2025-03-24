package datastruct

import (
    "log"
)

func NewUnion(v int) *Union {
    base := make([]int, v)
    for i := 0; i < v; i++ {
        base[i] = i
    }
    return &Union{
        total: v,
        base: base,
    }
}

type Union struct {
    total int
    base []int
}

func (u *Union) Find(v int) int {
    if v >= len(u.base) {
        log.Fatalln("illegal number,great than the max number")
    }
    queue := make([]int, 0)
    tmp := u.base[v]
    for tmp != u.base[tmp] {
        tmp = u.base[tmp]
        queue = append(queue, tmp)
    }
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        u.base[node] = tmp
    }
    return tmp
}

func (u *Union) Merge(a, b int) {
    f1 := u.Find(a)
    f2 := u.Find(b)
    if f1 == f2 {
        return
    }
    u.base[f1] = f2
    u.total--
}