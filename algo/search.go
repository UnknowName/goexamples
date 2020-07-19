package algo

func SeqSearch(nums []int, value int) int {
	for i, v := range nums {
		if v == value {
			return i
		}
	}
	return -1
}

// 非递归
func BinarySearch(nums []int, value int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		middle := (min + max) / 2
		if nums[middle] == value {
			return middle
		} else if nums[middle] > value {
			max = middle - 1
		} else if nums[middle] < value {
			min = middle + 1
		}
	}
	return -1
}

// 递归
func BinarySearchR(nums[]int, value, min, max int) int {
	if min > max {
		return -1
	}
	middle := (min + max) / 2
	if nums[middle] == value {
		return middle
	} else if nums[middle] > value {
		return BinarySearchR(nums, value, min, middle - 1)
	} else if nums[middle] < value {
		return BinarySearchR(nums, value, middle + 1, max)
	}
	return -1
}

/*
func OtherQuick(list []int, start, end int) {
	// 只剩一个元素时就返回了
	if start >= end {
		return
	}
	// 标记最左侧元素作为参考
	tmp := list[start]
	// 两个游标分别从两端相向移动，寻找合适的"支点"
	left := start
	right := end
	for left != right {
		// 右边的游标向左移动，直到找到比参考的元素值小的
		for list[right] >= tmp && left < right {
			right--
		}
		// 左侧游标向右移动，直到找到比参考元素值大的
		for list[left] <= tmp && left < right {
			left++
		}

		// 如果找到的两个游标位置不统一，就游标位置元素的值，并继续下一轮寻找
		// 此时交换的左右位置的值，右侧一定不大于左侧。可能相等但也会交换位置，所以才叫不稳定的排序算法
		if left < right {
			list[left], list[right] = list[right], list[left]
		}
		fmt.Println(list)
	}

	// 这时的left位置已经是我们要找的支点了，交换位置
	// fmt.Println(left, start)
	list[start], list[left] = list[left], tmp

	// 按支点位置吧原数列分成两段，再各自逐步缩小范围排序
	OtherQuick(list, start, left - 1)
	OtherQuick(list, left + 1, end)
}
*/