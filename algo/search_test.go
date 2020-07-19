package algo

import (
	"fmt"
	"testing"
)


func TestSeqSearch(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7}
	index := SeqSearch(nums, 7)
	fmt.Println(index)
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7,8,11}
	index := BinarySearch(nums, 11)
	fmt.Println(index)
}

func TestBinarySearchR(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7,8,11}
	index := BinarySearchR(nums, 10, 0, len(nums) - 1)
	fmt.Println(index)
}

/*
func TestOtherQuick(t *testing.T) {
	nums := []int{45, 38, 66, 90, 88, 10, 25, 45}
	// nums := []int{8, 9, 1, 7, 2, -100, 200, 3, 5, 4, 6, 0}
	OtherQuick(nums, 0, len(nums) - 1)
	fmt.Println(nums)
}
 */
