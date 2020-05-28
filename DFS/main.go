package main

import (
	"fmt"

	"../Graph"
	rbt "../RedBlackTree"
)

type DepFirstSearch interface {
	DFS(s *Graph.Node)
}

type DFSResult struct {
	//索引表示node的key值 值表示
	Marked *rbt.RBTree
	Count  int
}

func FindDFS(g *Graph.Graph, node *Graph.Node) *DFSResult {
	if node.InGraph != g {
		fmt.Printf("Node is not in this graph.")
		return nil
	}
	d := new(DFSResult)
	d.Marked = rbt.NewRBTree()
	rbnode := rbt.NewRBNode(node.Key, node)
	// fmt.Printf("%p\n", rbnode)
	d.Marked.Put(rbnode)
	d.DFS(node)
	return d
}

func (d *DFSResult) DFS(s *Graph.Node) {
	// fmt.Printf("查看%d的路徑 ", s.Key)
	// if len(s.Adjs) > 0 {
	// 	fmt.Printf("鄰接點有")
	// 	for _, v := range s.Adjs {
	// 		fmt.Printf("%d ", v.Key)
	// 	}
	// 	fmt.Println("")
	// }
	for _, node := range s.Adjs {
		rbnode := rbt.NewRBNode(node.Key, node)
		if d.Marked.Find(rbnode) == nil {
			d.Marked.Put(rbnode)
			d.DFS(node)
		}
	}
	d.Count++
}

// func (d *DFSResult)IsMarked(s *Graph.Node) bool{
// 	if d.Marked.Find()
// }
func (d *DFSResult) ShowSearch() {
	fmt.Println("透過紅黑樹儲存路徑: 相通的點如下:")
	d.Marked.Show()
}

func main() {
	g := Graph.NewGraph()
	// var g.Nodes [100]*Graph.Node
	for i := 0; i < 100; i++ {
		node := Graph.NewNode(i)
		g.AddNode(node)
	}
	g.AddEdge(g.Nodes[0], g.Nodes[1])
	g.AddEdge(g.Nodes[0], g.Nodes[2])
	g.AddEdge(g.Nodes[0], g.Nodes[3])
	g.AddEdge(g.Nodes[3], g.Nodes[5])
	g.AddEdge(g.Nodes[5], g.Nodes[10])
	g.AddEdge(g.Nodes[90], g.Nodes[99])
	g.AddEdge(g.Nodes[99], g.Nodes[80])
	g.AddEdge(g.Nodes[50], g.Nodes[51])
	result := FindDFS(g, g.Nodes[90])
	result.ShowSearch()

}
