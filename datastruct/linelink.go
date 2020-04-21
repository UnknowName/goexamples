package datastruct

import (
	"errors"
	"fmt"
)

const MAXSIZE = 10

// 线性表的中数据元素
type Element struct {
	// 数据项
	Name  string
	ID   int
}

// 线性表
type LineLink struct {
	length int
	Data   [MAXSIZE]Element
}

// 线性表创建
func NewLineLink() *LineLink {
	return &LineLink{}
}

// 插入元素。如果多次在相同位置插入，最后插入的将会覆盖老的数据
func (l *LineLink) Insert(element Element, i int) error {
	//检查是否已满，已满返回线性表已满
	if l.length >= MAXSIZE {
		fmt.Println("线性表已满")
		return errors.New("线性表已满")
	}
	// 检查插入值i的范围是否合法
	// 因为索引是从0开始，所以最后要加1
	if (i < 0 ) || (i > l.length + 1) {
		fmt.Println("插入位置不合法")
		return errors.New("插入的位置不合法")
	}
	// 将已存在元素逐个向后移动， 为i腾出位置
	for j := l.length; j >= i; j-- {
		l.Data[j] = l.Data[j - 1]
	}
	l.Data[i - 1] = element
	// 插入成功后，长度加1
	l.length += 1
	return nil
}

// 删除索引中的元素.因为索引从0开始，而判断是从1开始，所以i-1的位置之后的元素都要向前移动
func (l *LineLink) Delete(i int) error {
	if (i < 0) || (i > l.length + 1) {
		fmt.Println("i值不合法")
		return errors.New("i值不合法")
	}
	for j := i - 1; j <= l.length - 1; j++ {
		l.Data[j] = l.Data[j+1]
	}
	l.length -= 1
	return nil
}

// 返回element在线性表中的索引位置。索引从0开始
func (l *LineLink) Location(e Element) int {
	// 循环遍历表中的所有数据项
	for i, element := range l.Data {
		if element == e {
			return i
		}
	}
	return -1
}
