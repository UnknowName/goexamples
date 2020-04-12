package structs

import (
	"encoding/json"
	"fmt"
)

// 标准的Struct
type Version struct {
	Number string       `json:"number"`
	BuildHash string    `json:"build_hash"`
}

func (v *Version) GetVersion() {
	fmt.Println("My Version is ", v.Number)
}

type Elastic struct {
	Name string         `json:"name"`
	Tag  string         `json:"tagline"`
	// 非匿名嵌入，只有属性会继承，结构体的方法不会
	Version Version     `json:"version"`
}

func (e *Elastic) Echo() {
	fmt.Println("Nothing to do, echo")
}

func (e *Elastic) ToJson() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

// 自定义的类型增加方法
type Header map[string][]string

// 因为map本身就是一个引用类型，所以这里不使用指针，实际上它本身就是一个指针
func (h Header) Get(key string) string {
	if values := h[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

// 匿名结构体内嵌时，结构体的方法也会继承下来。非匿名不可以
type Another struct {
	// 匿名内嵌进来，同时方法也继承了下来
	Version
}

func (a *Another) Echo() {
	fmt.Println("Call Another echo method")
}