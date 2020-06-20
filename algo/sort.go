package algo

import "fmt"

/*
	思路: 每进行一轮循环，就确定最大值。同旁边的数据两两进行比较，如果大于，就交换位置
*/
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

/*
	思路:  每进行一轮循环，确定最小值，确定后该值进行对调。
	这样，每进行一轮就确定了一个最小值。
*/
func SelectSort(nums []int) {
	length := len(nums)
	var minIndex = 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[minIndex] > nums[j] {
				nums[minIndex], nums[j] = nums[j], nums[minIndex]
			}
		}
		minIndex++
	}
}

/*
	分成已排序与未排序两组数据，如果比第一个数大，不管它
	如果比第一个数小，将第一个数往后退一步，将第二个数插入第一个数的位置
 */
func InsertSort(nums []int) {
	length := len(nums)
	//第一次时，就从0的位置认为有序，遍历剩下的元素，同已排序的数据进行比较
	for i := 1; i < length; i++ {
		// 遍历已排序的数据.当i=1时，nums[:1]是认为有序的， 当i=2时，nums[:2]是认为有序的
		insertValue := nums[i]
		j := i - 1
		// 从右向左比较元素，当比较的元素小于已排好序的元素时，时移动元素
		for j >= 0 && insertValue < nums[j] {
			// 这里相当于前面的元素全部往后移动一位
			//[1, 0, 3, 4] => [1,1,3,4]
			nums[j+1] = nums[j]
			j--
		}
		// 上面的循环执行完成后，就找到了insertValue插入的位置
		nums[j + 1] = insertValue
	}
}

/*
	将大的数组，先分成length/2小组，小组排序成再在原来的基础上分length/2组，再排序。直到最后为一组，再排序。
	效率不好
 */
func ShellSortBySwap(nums []int) {
	count := len(nums) / 2
	// 第一步，对原始数组进行分组
	for step := count; step > 0; step = step / 2 {
		// 遍历数组中的所有元素，让每个分组中的第一个数与nums[i+step}的元素进行比较大小
		for i := 0; i < len(nums); i++ {
			// 同与步长的元素逐个进行比较，并交换
			for j := i + step; j < len(nums); j += step {
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
}


func ShellSortByInsert(nums []int) {
	length := len(nums)
	for step := length / 2; step > 0; step = step / 2 {
		for i := step; i < length; i++ {
			j := i
			temp := nums[j]
			if nums[j] < nums[j - step] {
				for j - step >=0 && temp < nums[j - step] {
					// 因为已分组，所以将原来的值移动到分组的索引，通过step定位索引
					nums[j] = nums[j - step]
					j -= step
				}
				nums[j] = temp
			}
		}
		fmt.Println("start ", nums)
	}
}


func QuickSort(nums []int) {
	nums = []int{45, 38, 66, 90, 88, 10, 25, 45}
	length := len(nums)
	left := 0
	right := length - 1
	// 第一次快速排序
	middle := nums[left]
	for {
		if right <= left {
			fmt.Println(left, right)
			postion := left
			for postion > 0 {
				nums[left - 1] = nums[left]
			}

			break
		}
		for nums[right] >= middle {
			right--
		}
		// temp := nums[right]
		for nums[left] <= middle {
			left++
		}
		nums[right], nums[left] = nums[left], nums[right]
		fmt.Println(nums)
		left++
		right--
		//break
	}
	fmt.Println(nums)
}