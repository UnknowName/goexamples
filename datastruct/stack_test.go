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
		// stack.Pop()
		v := stack.Top()
		fmt.Println(v)
	}
}

// 使用栈实现表达式求值 3 * 6 + 2 - 3，不支持括号
func TestStack_Pop(t *testing.T) {
	expr := "20/2/2-10+11"
	numberStack := NewStack(10)
	symbolStack := NewStack(10)
	preRune := 'S'
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
					symbolStack.Push(preSymbol)
					symbolStack.Push(v)
				}
			}
		} else {
			// 为数字时，直接入栈
			// 遍历时为rune类型, 原栈为interface{}类型
			// 检查上一个类型是不是数字
			if preRune == '+' || preRune == '-' || preRune == '*' || preRune == '/' || preRune == 'S' {
				// 说明前面没有数字，直接入栈
				num, _ := strconv.Atoi(string(v))
				numberStack.Push(num)
			} else {
				// 说明前面是一个数字，需要拼接
				preNum := numberStack.Pop()
				// ASCII码中，因为前一个字符类型为rune。数字是从48开始编码的
				_num := fmt.Sprintf("%d%d", preNum, v-48)
				num, _ := strconv.Atoi(_num)
				numberStack.Push(num)
			}
		}
		preRune = v
	}
	// 相关元素压入栈中再出栈进行计算
	// 如果连续出栈的符号是减或者是除,则将连续的取反后再进行计算如(1+2-3-1)
	result := 0
	for {
		num1 := numberStack.Pop()
		if numberStack.Empty() {
			fmt.Println(result)
			break
		}
		num2 := numberStack.Pop()
		_symbol := symbolStack.Pop()
		// 接着取下一个运算符，看是不是同上一个一样为-或者/，是的话计算取反后压入栈中
		var nextSymbol rune
		if !symbolStack.Empty() {
			_nextSymbol := symbolStack.Pop()
			nextSymbol = _nextSymbol.(rune)
		}
		symbol := _symbol.(rune)
		if nextSymbol == symbol && nextSymbol == '-' {
			result = calcul(num1.(int), num2.(int), '+')
			symbolStack.Push(nextSymbol)
			numberStack.Push(result)
		} else if nextSymbol == symbol && symbol == '/' {
			result = calcul(num1.(int), num2.(int), '*')
			symbolStack.Push(nextSymbol)
			numberStack.Push(result)
		} else {
			result = calcul(num1.(int), num2.(int), symbol)
			numberStack.Push(result)
		}
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

// 使用栈实现中缀表达式转为后缀表达式
/*
	思路:
	1. 遇到数字，直接输出
	2. 遇到运算符，如果栈为空或者栈顶元素为左括号，直拉入栈
	3. 遇到括号，如果是左括号，直接入栈
		3.1 如果遇到右括号，则将栈中的运算符一直出栈并输出，直到遇到左括号，左括号不输出
		3.2 遇到其他运算符从栈中弹出元素直到遇到发现更低优先级的元素(或者栈为空)为止或者遇到左括号。
			弹出完这些元素后，再将遇到的操作符压入到栈中。
		3.3 有一点需要注意，只有在遇到" ) "的情况下我们才弹出" ( "，其他情况我们都不会弹出" ( "。
 */
func TestStack_Get(t *testing.T) {
	express := "1+2*3+(4*5+6)*9"
	var afterExpress string
	symbolStack := NewStack(10)
	for _, v := range express {
		// 为数字直接入栈
		if v >= '0' && v <= '9' {
			afterExpress += string(v)
		} else {
			symbol := string(v)
			// 2 如果栈为空或者栈顶为左括号，或者待入栈为左括号，直接入栈
			if symbol == "(" || symbolStack.Empty() || symbolStack.Top().(string) == "(" {
				symbolStack.Push(symbol)
			} else if symbol == ")" {
				// 为右括号，栈中必须有元素，如果没有元素说明程序前段有误或者表达式不对
				// 为右括号情况,需要将当前栈内的元素弹出并输出，直到遇到左括号
				preSymbol := symbolStack.Pop().(string)
				if symbolStack.Top().(string) == "(" {
					// 左括号弹出，但不输出
					symbolStack.Pop()
				}
				afterExpress += preSymbol
			} else {
				// 其他情况，一直弹出直到遇到比symbol低优先级或者或者左括号
				// 1. 比较优先级，如果栈顶的优先级低于当前或者为空，直接入栈
				if low(symbol, symbolStack.Top().(string)) {
					symbolStack.Push(symbol)
				} else {
					// 栈顶优先级大于等于待入元素，一直弹出，直到优先级小，栈为空，栈顶为左括号为止
					// 弹出后再将当前元素入栈
					for {
						if symbolStack.Empty() || symbolStack.Top().(string) == "(" ||
							low(symbol,symbolStack.Top().(string)){
							break
						}
						preSymbol := symbolStack.Pop().(string)
						afterExpress += preSymbol
					}
					// 弹出后再压入栈中
					symbolStack.Push(symbol)
				}
			}
		}
	}
	// 最后将栈中的运算符输出
	for {
		if symbolStack.Empty() {
			break
		}
		symbol := symbolStack.Pop().(string)
		afterExpress += symbol
	}
	fmt.Println(afterExpress)
}


func low(v1, v2 string) bool {
	value1 := 0
	value2 := 0
	if v1 == "*" || v1 == "/" {
		value1 = 10
	}
	if v2 == "*" || v2 == "/" {
		value2 = 10
	}
	return value2 < value1
}