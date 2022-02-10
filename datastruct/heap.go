package datastruct

import "fmt"

func NewHeap(nums []int) *Heap {
    heap := &Heap{values: nums, size: len(nums)}
    for i := (len(nums) - 1) / 2; i >= 0; i-- {
        heap.moveDown(i)
    }
    return heap
}

type Heap struct {
    values []int
    size   int
}

func (h *Heap) Size() int {
    return h.size
}

func (h *Heap) String() string {
    return fmt.Sprintf("Heap{%v}", h.values)
}

func (h *Heap) Push(v int) {
    h.values = append(h.values, v)
    h.moveUp(h.size)
    h.size += 1
}

func (h *Heap) Pop() int {
    if h.size >= 1 {
        tmp := h.values[0]
        h.values[0] = h.values[h.size - 1]
        h.size -= 1
        h.values = h.values[:h.size]
        h.moveDown(0)
        return tmp
    }
    return 0
}

// 同自身左右孩子比较，并决定是否要交换顺序
func (h *Heap) moveDown(index int) {
    maxIndex := index
    // 存在左孩子
    for index * 2 + 1 < h.size {
        // 左孩子大于自己，将索引修改为左孩子的索引
        if h.values[index * 2 + 1] > h.values[index] {
            maxIndex = index * 2 + 1
        }
        // 有右孩子且右孩子大于最大索引上的值，修改最大索引为右孩子
        if index * 2 + 2 < h.size && h.values[maxIndex] < h.values[index * 2 + 2] {
            maxIndex = index * 2 + 2
        }
        // 如果当前最大值是自己，中止循环
        if maxIndex == index {
            break
        }
        // 交换
        h.values[index], h.values[maxIndex] = h.values[maxIndex], h.values[index]
        // 更新索引为最大值的索引
        index = maxIndex
    }
}

// 同父节点比较，是不是满足条件
func (h *Heap) moveUp(index int) {
    for (index - 1) / 2 >= 0 {
        parentIndex := (index - 1) / 2
        if h.values[index] > h.values[parentIndex] {
            h.values[index], h.values[parentIndex] = h.values[parentIndex], h.values[index]
        }
        index = parentIndex
        // 当索引为0时，跳出循环
        if index == 0 {
            break
        }
    }
}


