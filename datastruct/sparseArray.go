package datastruct

import "fmt"

func SparseArray() {
	fmt.Println("hello, struct")

	array := [][]int {
		{0, 0, 0, 1},
		{1, 0, 0, 4},
		{0, 0, 0, 0},
		{5, 0, 0, 0},
		{0, 0, 0, 0},
	}
	fmt.Println("origin array ", array)
	sparse := array2Sparse(array)
	fmt.Println("Array to sparse array ", sparse)
	originArray := sparse2Array(sparse)
	fmt.Println("Sparse array to array ", originArray)
}

func array2Sparse(array [][]int) [][3]int {
	var sum int
	for _, row := range array {
		for _, v := range row {
			if v != 0 {
				sum ++
			}
		}
	}
	// 初始化稀疏数组，因为第一行要存放基础信息，所以长度为有效数据+1
	sparse := make([][3]int, sum + 1)
	for i := range sparse {
		sparse[i] = [3]int{}
	}
	// 第一行存入原始数组的行与列以及有效数字
	sparse[0][0] = len(array)
	sparse[0][1] = len(array[0])
	sparse[0][2] = sum
	// 再次循环，将有效数字存入稀疏数组中
	var count = 1
	for i, row := range array {
		for j, v := range row {
			if v != 0 {
				sparse[count][0] = i
				sparse[count][1] = j
				sparse[count][2] = v
				count++
			}
		}
	}
	return sparse
}

func sparse2Array(sparse [][3]int) [][]int {
	row := sparse[0][0]
	col := sparse[0][1]
	validValue := sparse[0][2]
	array := make([][]int, row)
	for i := range array {
		array[i] = make([]int, col)
	}
	// 第一行存储了额外信息，所以要忽略掉第一行
	for i := 1; i <= validValue; i ++ {
		valueRow := sparse[i][0]
		valueCol := sparse[i][1]
		value := sparse[i][2]
		array[valueRow][valueCol] = value
	}
	return array
}