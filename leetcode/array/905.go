package array

func SortArrayByParity2(nums []int) []int {
    // 变成奇偶奇偶这样的
    // 双指针互换
    i := 0
    j := len(nums) - 1
    for i < j {
        for i < j && i % 2 == 0 && nums[i] % 2 != 0 {
            i++
        }
        for i < j && j % 2 != 0 && nums[j] % 2 == 0 {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
    return nums
}


func SortArrayByParity(nums []int) []int {
    // 直接奇数放后面，偶数放前面
    i := 0
    j := len(nums) - 1
    for i < j {
        if i < j && nums[i] % 2 == 0 {
            i++
            continue
        }
        if i < j && nums[j] % 2 != 0 {
            j--
            continue
        }
        if i < j && nums[i] % 2 != 0 && nums[j] % 2 == 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        }
    }
    return nums
}
