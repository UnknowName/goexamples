package algo

func BubbleSort(nums []int) {
	length := len(nums)
	flag := false
	for j := 0; j < length; j++ {
		// 这所以要减j,是每循环一轮，相应倒数第几的数是确定下来的
		for i := 0; i < length-1-j; i++ {
			if nums[i] > nums[i+1] {
				flag = true
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		if !flag {
			break
		} else {
			flag = false
		}
	}
}
