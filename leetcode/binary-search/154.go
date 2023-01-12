package binary_search

func FindMin(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        pivot := low + (high - low) / 2
        // 中间值小于最右边的值
        if nums[pivot] < nums[high] {
            // 在左边查找
            high = pivot
        } else if nums[pivot] > nums[high] {
            // 中间值大于最右边的值，在右边查找
            low = pivot + 1
        } else {
            // 等于的情况，将数组慢慢减少搜索范围
            high--
        }
    }
    return nums[low]
}