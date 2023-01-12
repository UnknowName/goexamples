package algo

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	lst := []int{-9, 1, 10, -1, 1, 5, 30}
	BubbleSort(lst)
	fmt.Println(lst)
}

func TestSelectSort(t *testing.T) {
	// lst := []int{-9, 20, -2, 1, 90}
	lst2 := []int{-22,0, -1, 2, -9, -10, 8, 7}
	// SelectSort(lst)
	SelectSort(lst2)
	fmt.Println(lst2)
}

func TestInsertSort(t *testing.T) {
	lst := []int{1, 2, -10, 0, -20, 8}
	InsertSort(lst)
	fmt.Println(lst)
}

func TestShellSortBySwap(t *testing.T) {
	lst := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	ShellSortBySwap(lst)
	fmt.Println(lst)
}

func TestShellSortByInsert(t *testing.T) {
	lst := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	ShellSortByInsert(lst)
	fmt.Println(lst)
}

func TestQuickSort(t *testing.T) {
	// lst := []int{45, 38, 66, 90, 88, 10, 25, 45}
	lst := []int{8, 9, 1, 7, 2, -100, 200, 3, 5, 4, 6, 0}
	QuickSort(lst, 0, len(lst) - 1)
	fmt.Println(lst)
}

func TestMergeSort(t *testing.T) {
	lst := []int{8, 9, 1, 7, 2, -100, 200, 3, 5, 4, 6, 0,11}
	nums := MergeSort(lst)
	fmt.Println(nums)
}

func TestMerge(t *testing.T) {
	nums1 := []int{0,1,2,3}
	num2 := []int{2,5,6}
	result := merge(nums1, num2)
	fmt.Println(result)
}

func TestBucketSort(t *testing.T) {
	lst := []int{8, 9, 1, 7, 2, -100, 200, 3, 5, 4, 6, 0,11}
	BucketSort(lst)
	fmt.Println(lst)
}