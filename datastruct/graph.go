package datastruct

import (
    "log"
)

func NewGraph(maxNode, maxEdge int) *Graph {
	return &Graph{
	    nodes: make([]int, maxNode),
	    nextEdges: make([]int, maxEdge),
	    toNodes: make([]int, maxEdge),
	    inbounds: make([]int, maxNode),
	    edgeNum: 0,
    }
}

type Graph struct {
	nodes     []int // 图节点的初始边，为0表示没有从该节点出发的边
	nextEdges []int // 边的下一条边
	toNodes   []int // 边去往的节点
	inbounds  []int // 节点的入度表
	edgeNum   int   // 边号
}

func (g *Graph) AddEdge(from, to int) {
    if from >= len(g.nodes) || to >= len(g.nodes) {
        log.Fatalln("illegal data,the from or to can't >=", len(g.nodes))
    }
    g.edgeNum++
    g.inbounds[to]++
    g.nextEdges[g.edgeNum] = g.nodes[from]
    g.nodes[from] = g.edgeNum
    g.toNodes[g.edgeNum] = to
}

func (g *Graph) TopoSort() []int {
    queue := make([]int, 0)
    for node := range g.nodes {
        if g.inbounds[node] == 0 {
            queue = append(queue, node)
        }
    }
    ans := make([]int, 0)
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for edge := g.nodes[node]; edge > 0; edge = g.nextEdges[edge] {
            toNode := g.toNodes[edge]
            g.inbounds[toNode]--
            if g.inbounds[toNode] == 0 {
                queue = append(queue, toNode)
            }
        }
        ans = append(ans, node)
    }
    if len(ans) != len(g.nodes) {
        return nil
    }
    return ans
}