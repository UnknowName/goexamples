package design_pattern

import "testing"

func TestNewAreaVisitor(t *testing.T) {
    c := &circle{}
    // 计算面积
    areaVisitor := NewAreaVisitor()
    c.accept(areaVisitor)
    // 计算边
    edgeVisitor := NewEdgeVisitor()
    c.accept(edgeVisitor)
}
