package datastruct

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// stack.Push("a")
	for i := 0; i < 5; i++ {
		v := stack.Pop()
		fmt.Println(v)
	}
}

// 使用栈实现表达式求值 3 * 6 + 2 - 3，不支持括号
func TestStack_Pop(t *testing.T) {
	expr := "2+2*4"
	numberStack := NewStack(10)
	symbolStack := NewStack(10)
	for _, v := range expr {
		// 先判断轮询的字符为数字还是符号
		if v == '+' || v == '-' || v == '/' || v == '*' {
			// 运算符入栈逻辑
			if symbolStack.Empty() {
				// 为空直接入栈
				symbolStack.Push(v)
			} else {
				/*
					不为空时，判断栈内的符号优先级。如果当前操作符优先级小于等于符号栈中的操作符，
					从数字栈中取两个数据，并从操作符栈中取一个操作符，计算后的结果再入数字栈
					注意点: 减法与除法是后面的数字操作前面的。
				*/
				preSymbol := symbolStack.Pop()
				if priority(v, preSymbol.(rune)) {
					// true为小于等于。数字栈出两个，符号栈出一个.计算后将结果再入栈
					_num1 := numberStack.Pop()
					_num2 := numberStack.Pop()
					symbol := preSymbol.(rune)
					sum := calcul(_num1.(int), _num2.(int), symbol)
					// 原来的符号入栈
					numberStack.Push(sum)
					symbolStack.Push(v)
				} else {
					// 大于时除了当前的符号要入栈， 前面取出的符号也要入栈
					//TODO 问题出在这里。当下一个运行符大于栈内时，有问题
					symbolStack.Push(preSymbol)
					symbolStack.Push(v)
				}


			}
		} else {
			// 为数字时，直接入栈
			// 遍历时为rune类型, 原栈为interface{}类型
			num, _ := strconv.Atoi(string(v))
			numberStack.Push(num)
		}
	}
	// 遍历完计算一次
	// symbolStack.Show()
	// numberStack.Show()
	result := 0
	for {
		num1 := numberStack.Pop()
		if numberStack.Empty() {
			fmt.Println(result)
			break
		}
		// num1 := numberStack.Pop()
		num2 := numberStack.Pop()
		_symbol := symbolStack.Pop()
		symbol := _symbol.(rune)
		result = calcul(num1.(int), num2.(int), symbol)
		numberStack.Push(result)
	}

}

// 判断数字
// 优先低为true,高为false
func priority(v1, v2 rune) bool {
	value1 := 0
	value2 := 0
	if v1 == '*' || v1 == '/' {
		value1 = 10
	}
	if v2 == '*' || v2 == '/' {
		value2 = 10
	}
	return value1 <= value2
}

func calcul(num1 int, num2 int, operator rune) int {
	sum := 0
	switch operator {
	case '+':
		sum = num2 + num1
	case '-':
		sum = num2 - num1
	case '/':
		sum = num2 / num1
	case '*':
		sum = num2 * num1
	}
	return sum
}

func TestStack_Push(t *testing.T) {
	v1 := '/'
	v2 := '+'
	fmt.Println(priority(v1, v2))
}