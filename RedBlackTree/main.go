package RedBlackTree

// RedBlackTree

import (
	"fmt"
	"os"
)

type RBTree struct {
	Head *RBNode
	//節點中的元素數
	N int
}

func NewRBTree() *RBTree {
	tree := new(RBTree)
	tree.N = 0
	return tree
}

func (tree *RBTree) Put(newnode *RBNode) *RBTree {
	tree.Head = tree.Head.PutNode(newnode)
	//新增一顆新樹
	newtree := NewRBTree()
	newtree.Head = newnode
	tree.N++
	return newtree
}

func (n *RBTree) Find(node *RBNode) *RBNode {
	if n.Head != nil {
		h := n.Head.FindNode(node)
		if h != nil {
			return h
		}
	}
	return nil
}

type RBNode struct {
	Key   int
	Value interface{}
	Left  *RBNode
	Right *RBNode
	//true 代表紅色 false is black
	Color bool
}

//NewRBNode make a new RBNode
func NewRBNode(key int, val interface{}) *RBNode {
	n := new(RBNode)
	n.Key = key
	n.Value = val
	return n
}

func (n *RBNode) FindNode(node *RBNode) (fnode *RBNode) {
	if n == nil {
		return
	}
	if n.Key < node.Key {
		return n.Right.FindNode(node)
	} else if node.Key < n.Key {
		return n.Left.FindNode(node)
	} else {
		return n
	}
}
func (n *RBNode) PutNode(newnode *RBNode) *RBNode {
	// fmt.Println(n)
	if n == nil {
		newnode.Color = true
		return newnode
	}

	//比較鍵值
	if newnode.Key < n.Key {
		// fmt.Println("加入左邊")
		n.Left = n.Left.PutNode(newnode)
	} else if n.Key < newnode.Key {
		// fmt.Println("加入右邊")
		n.Right = n.Right.PutNode(newnode)
	} else {
		n.Value = newnode.Value
		return n
	}

	//判斷是否要左旋 當左子結點顏色為黑色 右子結點顏色為紅色
	//左邊可能為空
	if n.Left != nil {
		if (!n.Left.IsRed()) && (n.Right.IsRed()) {
			// fmt.Println("要左旋", n.Key, "左", n.Left, "右", n.Right)
			return n.RotateLeft()
		}
	} else { //左邊可能為空
		if n.Right.IsRed() {
			// fmt.Println("要左旋", n.Key, "左", n.Left, "右", n.Right)
			return n.RotateLeft()
		}
	}

	//或是右旋
	if n.Left != nil {
		if n.Left.Left != nil {
			if (n.Left.IsRed()) && (n.Left.Left.IsRed()) {
				// fmt.Println("要右旋", n.Key, "左", n.Left, "右", n.Right)
				n = n.RotateRight()
				//右旋以後顏色需要反轉
				if (n.Left.IsRed()) && (n.Right.IsRed()) {
					// fmt.Println("顏色反轉囉", n.Key, "左", n.Left, "右", n.Right)
					n.ColorFlip()
					return n
				}
			}
			return n
		}
	}
	if n.Left != nil && n.Right != nil {
		if (n.Left.IsRed()) && (n.Right.IsRed()) {
			// fmt.Println("顏色反轉囉", n.Key, "左", n.Left, "右", n.Right)
			n.ColorFlip()
			return n
		}
	}

	return n
}

func (tree *RBTree) Size() int {
	return tree.N
}

func (n *RBNode) IsRed() bool {
	return n.Color
}

// RotateLeft 將左旋的Node輸入 返回新的子節點
func (n *RBNode) RotateLeft() *RBNode {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = true
	return x
}

func (n *RBNode) RotateRight() *RBNode {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = true
	return x
}

func (n *RBNode) InfixShow() {
	if n != nil {
		n.Left.InfixShow()
		fmt.Printf("%p,Key:%d Value=%v Color=%v 左:%p 右:%p\n", n, n.Key, n.Value, n.Color, n.Left, n.Right)
		n.Right.InfixShow()
	}
}
func (tree *RBTree) Show() {
	if tree.Head != nil {
		fmt.Println("紅黑樹的根是", tree.Head.Key)
	}
	tree.Head.InfixShow()
}

//ColorFlip 顏色反轉
func (n *RBNode) ColorFlip() {
	n.Right.Color = false
	n.Left.Color = false
	n.Color = true
}

func main() {
	tree := NewRBTree()

	var k string
	key := 0
	var name string

	for {
		fmt.Println(`
  =======水滸英雄表======
  1. 添加英雄
  2. 顯示英雄
  3. 查詢英雄
  4. 退出
  請輸入你的選擇:
  `)
		fmt.Scan(&k)
		switch k {
		case "1":
			fmt.Println("請輸入英雄的編號:")
			fmt.Scan(&key)
			fmt.Println("請輸入英雄名稱:")
			fmt.Scan(&name)
			h := NewRBNode(key, name)
			tree.Put(h)
		case "2":
			tree.Show()
		case "3":
			fmt.Println("請輸入要搜尋的英雄編號:")
			fmt.Scan(&key)
			hero := NewRBNode(key, "")
			hero = tree.Find(hero)
			fmt.Println("你找的英雄是:", hero.Value)
		case "4":
			os.Exit(1)
		}
	}

}
