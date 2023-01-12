package algo

import (
	"fmt"
)

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


/*
	思路: 左右双指针，key一般取第一个，先从右边找一个比key小的数，再从左边找一个比key大的数，然后交换。
    左右指针向对方靠拢，当左指针等于右指针时，一轮排序完成。左边小于基准数，右边大于基准数，再分别对左右两边排序
 */
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	key := nums[left]
	for i != j {
		// 先从右边找小于key的值,并交换
		for nums[j] >= key && j > i {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

		// 再从左边找大于key的值，并交换
		for nums[i] <= key && j > i {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	QuickSort(nums, left, i - 1)
	QuickSort(nums, i + 1, right)
}


func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := MergeSort(nums[0:i])
	right := MergeSort(nums[i:])
	// fmt.Println(left, right)
	result := merge(left, right)
	return result
}


func merge(nums1, nums2 []int) []int {
	length1, length2 := len(nums1), len(nums2)
	result := make([]int, length1 + length2)
	i, j, k := 0, 0, 0
	// 两个数组比对，小的先放进结果集
	for i < length1 && j < length2 {
		if nums1[i] < nums2[j] {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
		k++
	}
	// 比对后再将剩余的数加入数组中，两个都要遍历过一次
	for i < length1 {
		result[k] = nums1[i]
		i++
		k++
	}
	for j < length2 {
		result[k] = nums2[j]
		j++
		k++
	}
	return result
}

func BucketSort(nums []int) {
	// 1. 先获取最大值与最小值，确定桶的个数以及基数
	if len(nums) < 2 {
		return
	}
	min := nums[0]
	max := nums[1]
	if min > max {
		nums[0], nums[1] = nums[1], nums[0]
		min, max = max, min
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	base := (max - min) / len(nums) + 1
	cnt := (max - min) / base + 1
	buckets := make([][]int, cnt)
	for _, num := range nums {
		index := (num - min) / base
		buckets[index] = append(buckets[index], num)
	}
	for i := range buckets {
		BucketSort(buckets[i])
	}
	i := 0
	for _, _nums := range buckets {
		if len(_nums) == 0 {
			continue
		}
		for j := range _nums {
			nums[i] = _nums[j]
			i++
		}
	}
}