package datastruct

import (
    "fmt"
    "testing"
)

func TestGraph_TopoSort(t *testing.T) {
    graph := NewGraph(5, 20)
    graph.AddEdge(1,2)
    graph.AddEdge(2,3)
    graph.AddEdge(3,4)
    res := graph.TopoSort()
    fmt.Println(res)
}
