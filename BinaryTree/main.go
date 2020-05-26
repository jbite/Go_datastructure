package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

func NewHero(no int, name string) *Hero {
	h := new(Hero)
	h.No = no
	h.Name = name
	return h
}

//前序遍歷 先輸出根結點 再左子樹 再輸出右子樹
func (h *Hero) PreOrder() {
	if h != nil {
		fmt.Printf("no:%d name=%s\n", h.No, h.Name)
		h.Left.PreOrder()
		h.Right.PreOrder()
	}
}

//中序遍歷
func (h *Hero) InfixOrder() {
	if h != nil {
		h.Left.InfixOrder()
		fmt.Printf("no:%d name=%s\n", h.No, h.Name)
		h.Right.InfixOrder()
	}
}
func (h *Hero) PostOrder() {
	if h != nil {
		h.Left.PostOrder()
		h.Right.PostOrder()
		fmt.Printf("no:%d name=%s\n", h.No, h.Name)
	}
}
func main() {
	root := NewHero(1, "宋江")
	left1 := NewHero(2, "吳用")
	right1 := NewHero(3, "盧俊義")
	right2 := NewHero(4, "林沖")
	node10 := NewHero(10, "魯達")
	node12 := NewHero(12, "董平")

	root.Left = left1
	root.Right = right1
	right1.Right = right2
	left1.Left = node10
	left1.Right = node12
	fmt.Println("====前序遍歷=====")
	root.PreOrder()
	fmt.Println("====中序遍歷=====")
	root.InfixOrder()
	fmt.Println("====後序遍歷=====")
	root.PostOrder()

}
