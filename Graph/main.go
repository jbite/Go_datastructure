package Graph

import "fmt"

type Node struct {
	Key     int
	Adjs    []*Node
	InGraph *Graph
}

type Graph struct {
	//頂點數量
	V int
	//邊的數量
	E     int
	Nodes []*Node
}

func NewGraph() *Graph {
	g := new(Graph)
	return g
}

func NewNode(key int) *Node {
	n := new(Node)
	n.Key = key
	return n
}

func (g *Graph) CountV() int {
	return g.V
}
func (g *Graph) CountE() int {
	return g.E
}
func (g *Graph) AddNode(n *Node) {
	for _, v := range g.Nodes {
		if n.Key == v.Key { //如果節點值重複就更新
			v.Key = n.Key
			return
		}
	}
	g.Nodes = append(g.Nodes, n)
	n.InGraph = g
	g.V++
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node) {
	//先判斷點是否在地圖中
	if n1.InGraph == g && n2.InGraph == g {
		n1.Adjs = append(n1.Adjs, n2)
		n2.Adjs = append(n2.Adjs, n1)
		g.E++
	}
	if n1.InGraph != g {
		fmt.Println(n1.Key, "不在地圖中")
	}
	if n2.InGraph != g {
		fmt.Println(n2.Key, "不在地圖中")
	}
}
func (g *Graph) ShowGraph() {
	fmt.Println("圖中的節點有:")
	for _, v := range g.Nodes {
		fmt.Printf("%d ", v.Key)
		v.ShowNeighbor()
	}
}
func (n *Node) ShowNeighbor() {

	if n.Adjs == nil {
		fmt.Printf("%d沒有鄰居\n", n.Key)
		return
	}
	fmt.Printf("%d的鄰居有:[", n.Key)
	for _, v := range n.Adjs {
		fmt.Printf("%d", v.Key)
	}
	fmt.Println("]")
}
func main() {
	g := NewGraph()
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddEdge(n1, n2)
	g.ShowGraph()

}
