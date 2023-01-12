package try

import (
	"fmt"
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1,2,3,0,0}
	nums2 := []int{-1, 9}
	Merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{1,7,11,15}
	result := TwoSum2(nums, 22)
	fmt.Println(result)
}

func TestRotate(t *testing.T) {
	// Rotate([]int{1,2,3,4,5}, 22)
}

func TestArrayPairSum(t *testing.T) {
	nums := []int{2,2,1,1,1,2,2}
	result := MajorityElement(nums)
	fmt.Println(result)
}

func TestShuffle(t *testing.T) {
	nums := []int{1,2,3,4,4,3,2,1}
	Shuffle(nums, len(nums)/2)
}

func TestKidsWithCandies(t *testing.T) {
	nums := []int{2,3,5,1,3}
	KidsWithCandies(nums, 3)
}

func TestMissingNumber(t *testing.T) {
	MissingNumber([]int{1})
}

func TestAverage(t *testing.T) {
	Average([]int{1000,4000,2000,3000,6000})
}

func TestReplaceElements(t *testing.T) {
	nums := []int{1}
	ReplaceElements(nums)
}

func TestCheckIfExist(t *testing.T) {
	fmt.Println(CheckIfExist2([]int{-2,0,10,-19,4,6,-8}))
}

func TestFindMaxAverage(t *testing.T) {
	result := FindMaxAverage([]int{-10,10,20},  2)
	fmt.Println(result)
}

func TestMoveZeroes(t *testing.T) {
	MoveZeroes([]int{0,2,0,3,12})
}

func TestMajorityElement3(t *testing.T) {
	x := MajorityElement3([]int{3,2,2,3})
	fmt.Println(x)
}

func TestDominantIndex(t *testing.T) {
	is := DominantIndex([]int{0,1})
	fmt.Println(is)
}

func TestMinimumAbsDifference(t *testing.T) {
	MinimumAbsDifference([]int{1,3,4})
}

func TestSetZeroes(t *testing.T) {
	SetZeroes([][]int{[]int{0,1,2}, []int{0,0,5}, []int{5,6,7}})
}

func TestPivotIndex(t *testing.T) {
	i := PivotIndex([]int{1, 7, 3, 6, 5, 6})
	fmt.Println(i)
}

func TestRotate3(t *testing.T) {
	Rotate([]int{1,2,3,4,5,6,7,8,9}, 2)
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1,3,1,2,0,5}
	results := MaxSlidingWindow(nums, 3)
	log.Println("results=", results)
}