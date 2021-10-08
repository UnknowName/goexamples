package design_pattern

import "log"

/*
组合模式
适用场景
1. 需要实现树状结构
2. 如果希望客户端代码以相同的方式处理简单与复杂的元素
 */

// 组合接口

type component interface {
    print()
    add(component)
}

// 学校

func NewSchool(name string) *school {
    return &school{
        name: name,
        colleges: make([]component, 0),
    }
}

type school struct {
    name string
    colleges []component
}

func (s *school) print() {
    log.Println("school name ", s.name)
    for _, _com := range s.colleges {
        _com.print()
    }
}

func (s *school) add(c component) {
    s.colleges = append(s.colleges, c)
}

// 院系

func NewCollege(name string) *college {
    return &college{
        name: name,
        majors: make([]component, 0),
    }
}

type college struct {
    name string
    majors []component
}

func (c *college) print() {
    log.Println("college name ", c.name)
    for _, maj := range c.majors {
        maj.print()
    }
}

func (c *college) add(com component) {
    c.majors = append(c.majors, com)
}

// 专业

type major struct {
    name string
}

func (m *major) print() {
    log.Println("major name ", m.name)
}

func (m *major) add(c component) {
}