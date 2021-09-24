package design_pattern

import "log"

/*
访问者模式 访问者模式允许你在结构体中添加行为， 而又不会对结构体造成实际变更
场景:
1. 如果你需要对一个复杂对象结构 （例如对象树） 中的所有元素执行某些操作， 可使用访问者模式。
2. 可使用访问者模式来清理辅助行为的业务逻辑。
3. 当某个行为仅在类层次结构中的一些类中有意义， 而在其他类中没有意义时， 可使用该模式。
 */

// 形状接口，所有对象已实现的接口

type shape interface {
    accept(visitor)
    getType() string
}

// 访问者接口,因为Golang不支持重载，所以实现不同方法，这个接口需要自己实现并新增

type visitor interface {
    visitForSquare(*square)
    visitForCircle(*circle)
    visitForTriangle(*triangle)
}

// 假设square,circle,triangle这些结构体/类早就定义好了。那么只需要添加一个accept()方法即可

type square struct {}

func (s *square) accept(v visitor) {
    // 不同形状调用接口不同方法，因为golang不支持方法重载
    v.visitForSquare(s)
}

func (s *square) getType() string {
    return "square"
}

type circle struct {}

func (c *circle) accept(v visitor) {
    // 不同形状调用接口不同方法，因为golang不支持方法重载
    v.visitForCircle(c)
}

func (c *circle) getType() string {
    return "circle"
}

type triangle struct {}

func (t *triangle) accept(v visitor) {
    // 不同形状调用接口不同方法，因为golang不支持方法重载
    v.visitForTriangle(t)
}

func (t *triangle) getType() string {
    return "triangle"
}

// 定义一个具体的访问者，这里叫areaVisitor，计算面积

func NewAreaVisitor() *AreaVisitor {
    return &AreaVisitor{}
}

type AreaVisitor struct {
    area float64
}

func (a *AreaVisitor) visitForSquare(s *square) {
    log.Println("visitor for square", s)
}

func (a *AreaVisitor) visitForCircle(c *circle) {
    log.Println("visitor for circle", c)
}

func (a *AreaVisitor) visitForTriangle(t *triangle) {
    log.Println("visitor for triangle", t)
}

// 定义另一个具体的访问者，计算边

func NewEdgeVisitor() *EdgeVisitor {
    return &EdgeVisitor{}
}

type EdgeVisitor struct {

}

func (ev *EdgeVisitor) visitForSquare(s *square) {
    log.Println("calculate square edge ", s)
}

func (ev *EdgeVisitor) visitForCircle(c *circle) {
    log.Println("calculate circle edge ", c)
}

func (ev *EdgeVisitor) visitForTriangle(t *triangle) {
    log.Println("calculate triangle edge ", t)
}