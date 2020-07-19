package try

import (
	"fmt"
	"math"
	"sort"
)

// #88
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1与nums2有序，len(nums1) >= m + n
	// m与n分别代表当前数组的有效数据个数
	// 只能原地修改nums1
	// nums1 := []int{1,2,3,0,0}
	// nums2 := []int{2, 7}
	// 思路: 双指针，从后向前检查，大的放入后面
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	if i == -1 {
		for i := j; i >= 0; i-- {
			fmt.Println(i)
			nums1[i] = nums2[i]
		}
	}
}

// #167
// 因为有序，从小到大。所以可以使用双指针移动。无序时无效
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left != right {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		} else if numbers[left]+numbers[right] > target { // 大于目标值，则将右边最大值往左移，将相加数减少
			right--
		} else if numbers[left]+numbers[right] < target { // 小于目标值，向右移动，增加相关加数
			left++
		}
	}
	return []int{}
}

// 使用Map记录已遍历过的数字
func TwoSum2(numbers []int, target int) []int {
	var result []int
	dic := make(map[int]int)
	for i, v := range numbers {
		othNum := target - v
		if _, ok := dic[othNum]; ok {
			// 为什么首先要放dic里面的数据，因为dic里面的数据存放的是已遍历过的数据，肯定在前
			result = append(result, dic[othNum])
			result = append(result, i)
		}
		// 没找到，将值为KEY，索引为V放入字典中
		dic[v] = i
	}
	return result
}

// #561
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)
	var total int
	for i := 0; i < len(nums); i += 2 {
		total += nums[i]
	}
	return total
}

// #169 HashMap
func MajorityElement(nums []int) int {
	total := make(map[int]int)
	for _, v := range nums {
		if _, ok := total[v]; ok {
			total[v] += 1
		} else {
			total[v] = 1
		}
	}
	var result int
	var key int
	for k, v := range total {
		if v > result {
			result = v
			key = k
		}
	}
	return key
}

// #169 摩尔投票法。时间复杂度与空间复杂一样，没HashMap好理解
func MajorityElement2(nums []int) int {
	var tmp, cnt int
	for _, v := range nums {
		if cnt == 0 {
			tmp = v
			cnt++
		} else if tmp == v {
			cnt++
		} else {
			cnt--
		}
	}
	fmt.Println(tmp)
	return tmp
}

// #1470
func Shuffle(nums []int, n int) []int {
	result := make([]int, n*2)
	j := 0
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[j]
		result[i+1] = nums[j+n]
		j++
	}
	return result
}

// #1431
func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	// 先取出最大值,假设最大值的索引为0
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	// 将最大数减去其他的数，判断余数是不是小于等于extraCandies
	for i := 0; i < len(candies); i++ {
		result[i] = max-candies[i] <= extraCandies
	}
	return result
}

// #268 高斯求和公式 n(n+1)/ 2
func MissingNumber(nums []int) int {
	var sum int
	length := len(nums)
	sum = (length * (length + 1)) / 2
	for _, v := range nums {
		sum = sum - v
	}
	return sum
}

// #1491 注意不要丢失了精度
func Average(salary []int) float64 {
	var sum float64
	sort.Ints(salary)
	for _, v := range salary[1:len(salary)-1] {
		sum += float64(v)
	}
	return sum / float64(len(salary)-2)
}

// #1299
func ReplaceElements(arr []int) []int {
	// 从右往左遍历数组，先假设最后一个数字为最大值
	// 当后续数字大于它时，将最大值换掉
	length := len(arr)
	result := make([]int, length)
	if length == 1 {
		result[0] = -1
		return result
	}
	i := length - 1
	max := arr[length - 1]
	for i >= 0 {
		result[i] = max
		if arr[i] > max {
			max = arr[i]
		}
		i--
	}
	result[length-1] = -1
	// fmt.Println(result)
	return result
}

// #1346， 时间复杂度好大
func CheckIfExist(arr []int) bool {
	// 双指针方法.i与j = i+1. 先排序
	// 如果arr[i] * 2 > arr[j]，说明arr中不存在arr[i] * 2的数，这时可以让i++移动指针，继续找下一个
	// 如果arr[i] * 2 < arr[j],则继续将j++，直接j=len(arr) - 1后结束
	sort.Ints(arr)
	fmt.Println(arr)
	i := 0
	for i < len(arr) {
		j := i + 1
		for j < len(arr) {
			if arr[i] == 0 && arr[j] == 0 {
				return true
			}
			var total int
			if arr[i] < 0 {
				if arr[i] % 2 == 0 {
					total = arr[i] / 2
				} else {
					break
				}
			} else {
				total = arr[i] * 2
			}

			if total > arr[j] {
				j++
				continue
			} else if total < arr[j] {
				// 第一个元素 * 2都要小于下一个元素，说明该元素*2不存放于数组中，因为已排序
				// 跳出当前循环
				break
			} else {
				return true
			}
			j++
		}
		i++
	}
	return false
}

// #1346 使用HashMap
func CheckIfExist2(arr []int) bool {
	// 先将所有数据放入字典中
	dic := make(map[int]int)
	for _, v := range arr {
		if _, ok := dic[v]; ok {
			dic[v] += 1
		} else {
			dic[v] = 1
		}
	}

	for key, value := range dic {
		// fmt.Println(key, value)
		if key == 0 && value >= 2 {
			return true
		}
		for k, _ := range dic {
			var tmp int
			if k < 0 && k % 2 == 0 && k != 0 {
				tmp = k / 2
			} else {
				tmp = k * 2
			}
			if tmp == key && key != 0 {
				return true
			}
		}
	}
	return false
}

// #643 双重循环会超时
func FindMaxAverage(nums []int, k int) float64 {
	if k == 1 {
		sort.Ints(nums)
		return float64(nums[len(nums) - 1])
	}
	length := len(nums)
	var total float64
	for i := 0; i < k; i++ {
		total += float64(nums[i])
	}
	for j := 1; j < length; j++ {
		tmp := total - float64(nums[j - 1]) + float64(nums[k+1])
		fmt.Println(tmp)
	}
	fmt.Println(length, total)
	return total / float64(k)
}
/*
class Solution(object):
    def findMaxAverage(self, nums, k):
        if k == 1:
            return sorted(nums)[-1]

        sumnum = 0
        #初始值
        for i in range(k):
            sumnum += nums[i]
        maxnum = sumnum

        #每移动一次，减去头部的值，加上尾部的值
        for j in range(1, len(nums)-k+1):
            sumnum = sumnum - nums[j-1] + nums[j+k-1]
            maxnum = max(sumnum, maxnum)

        return float(maxnum) / k
 */

// # 283
func MoveZeroes(nums []int) {
	// 让非0同0进行交换，这样0就全部移动到了后面
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)
}

// # 1710
func MajorityElement3(nums []int) int {
	if len(nums) == 1 {
		return -1
	}
	count := make(map[int]int)
	for _, v := range nums {
		if _, ok := count[v]; ok {
			count[v] += 1
			if count[v] > len(nums) / 2 {
				return v
			}
		} else {
			count[v] = 1
		}
	}
	return -1
}

// # 747
func DominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// 先用最笨的办法，找出最大数，并记录索引值
	// 不能排序，排序索引就全乱了
	max := nums[0]
	maxIndex := 0
	i := 1
	for i < len(nums) {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
		i++
	}
	// fmt.Println(maxIndex, max)
	// 再用比较最大值是不是大于所有，大于就最终返回maxIndex,否则返回-1
	j := 0
	for j < len(nums) {
		if j == maxIndex || max >= nums[j] * 2 {
			j++
			continue
		} else {
			return -1
		}
	}
	return maxIndex
}

// # 53
func MaxSubArray(nums []int) int {
	total := nums[0]
	j := 1
	for j < len(nums) {
		if total + nums[j] > total {
			total += nums[j]
		}
		j++
	}
	fmt.Println(total)
	return total
}

// # 1200
func MinimumAbsDifference(arr []int) [][]int {
	var result [][]int
	if len(arr) == 2 {
		result = append(result, []int{arr[0], arr[1]})
		return result
	}
	sort.Ints(arr)
	min := math.Abs(float64(arr[0] - arr[1]))
	result = append(result, []int{arr[0], arr[1]})
	i, j := 1, 2
	for i < len(arr) && j < len(arr) {
		tmp := math.Abs(float64(arr[j] - arr[i]))
		if tmp < min {
			// 一旦小于，将之前的数据清空，将当前数据写入
			result = make([][]int,0)
			// 清空数据后，再将当前数据谢谢
			result = append(result, []int{arr[i], arr[j]})
			min = tmp
		} else if tmp == min {
			// 问题出在这里
			result = append(result, []int{arr[i], arr[j]})
		}
		i++
		j++
	}
	return result
}


// 0108 M* 的矩阵中有0，则所在的行与列则全部设置为0
func SetZeroes(matrix [][]int)  {
	// 先遍历矩阵，记录所有为0的索引
	// 这里用字典，因为只需要记录不重复的
	rowIndex := make(map[int]struct{})
	colIndex := make(map[int]struct{})
	for i, rows := range matrix {
		for j, v := range rows {
			if v == 0 {
				rowIndex[i] = struct{}{}
				colIndex[j] = struct{}{}
			}
		}
	}
	// fmt.Println(rowIndex, colIndex)
	// 将所在行有0的行全部置为0
	for row := range rowIndex {
		fmt.Println(row)
		rows := matrix[row]
		for i, _ := range rows {
			rows[i] = 0
		}
	}
	// 将所在列有0的列全部置为0
	for col := range colIndex {
		// fmt.Println(col)
		for k := 0; k < len(matrix);k++ {
			matrix[k][col] = 0
		}
	}
	// fmt.Println(matrix)
}

// #274
func PivotIndex(nums []int) int {
	// 1. 先计算数组中的总和
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	// fmt.Println(total)
	// 2. 当存在一个索引为左边的和等于右边的和时，leftSum 不断加，当到达该点时
	// 则 total - leftSum - x = leftSum， x所在的索引即为该值。否则不存在，最终返回-1
	var leftSum int
	for i, v := range nums {
		if total - v - leftSum == leftSum {
			return i
		}
		leftSum += v
	}
	return -1
}

// #189
func Rotate(nums []int, k int)  {
	length := len(nums)
	for i, v := range nums {
		fmt.Println(i, length - i - 1, v)
		nums[i], nums[length - i - 1] = nums[length - i - 1], nums[i]
	}
	fmt.Println(nums)
}

