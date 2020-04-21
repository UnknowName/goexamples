package interfaces

import (
	"fmt"
	"testing"
)

func TestChinese_Height(t *testing.T) {
	var chinese Humaner
	// Interface赋值时，是指针
	chinese = NewChinese("ZhangShan", 186)
	fmt.Println(chinese.Name())
	// 通过类型断言，将interface类型转换成struct类型
	c := chinese.(*Chinese)
	fmt.Printf("%T", c)
}
