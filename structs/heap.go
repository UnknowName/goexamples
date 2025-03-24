package structs

// 5个接口都要实现

type IntHeap []int

func (it IntHeap) Len() int {
    return len(it)
}

func (it IntHeap) Less(i, j int) bool {
    return it[i] > it[j]
}

func (it IntHeap) Swap(i, j int) {
    it[i], it[j] = it[j], it[i]
}

func (it *IntHeap) Push(v any) {
    *it = append(*it, v.(int))
}

// 从最后取堆

func (it *IntHeap) Pop() any {
    nums := *it
    v := nums[len(nums)-1]
    *it = nums[:len(nums)-1]
    return v
}
