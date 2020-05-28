package main

import "fmt"

//由內容表示組別 Index表示值
type UF struct {
	N    int
	Sets []*Element
}

type Element struct {
	Data   interface{}
	Parent *Element
}

func NewUF(n int) *UF {
	u := new(UF)
	u.N = n
	for i := 0; i < u.N; i++ {
		e := NewEle(i)
		u.Sets = append(u.Sets, e)
	}
	return u
}

func NewEle(data interface{}) *Element {
	ele := new(Element)
	ele.Data = data
	return ele
}

func (uf *UF) Union(a *Element, b *Element) {
	aRoot := uf.Find(a)
	bRoot := uf.Find(b)
	aRoot.Parent = bRoot
}

//搜尋元素所在的組別
func (uf *UF) Find(p *Element) *Element {
	for i := 0; i < len(uf.Sets); i++ {
		if p.Parent == uf.Sets[i] {
			return uf.Sets[i]
		}
	}
	return uf.Find(p.Parent)
}

func (uf *UF) AddToSet(ele *Element, n int) {
	ele.Parent = uf.Sets[n-1]
}
func main() {

	uf := NewUF(10)
	el1 := NewEle("我是小三")
	uf.AddToSet(el1, 3)
	el2 := NewEle("小四")
	uf.AddToSet(el2, 3)

	set := uf.Find(el2)
	fmt.Println(set)

}
