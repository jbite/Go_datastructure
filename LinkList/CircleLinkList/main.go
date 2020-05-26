package CircleLinkList

import (
	"errors"
	"fmt"
)

//定義一個環形鏈表
type CirLnLs struct {
	Index int
	Name  string
	Next  *CirLnLs
}

//工廠函數
func NewCirLnLs(index int, name string) *CirLnLs {
	c := new(CirLnLs)
	c.Index = index
	c.Name = name
	c.Next = c
	return c
}

func (c *CirLnLs) InsertNode(d *CirLnLs) {
	temp := c
	for {
		if temp.Next == c { //temp就是尾巴了
			temp.Next = d
			d.Next = c
			return
		}
		temp = temp.Next
	}
}

//DeleteNode return headnode, deleted node , error
func (c *CirLnLs) DeleteNode(index int) (temp *CirLnLs, delnode *CirLnLs, err error) {
	temp = c
	helper := c.NewHelper()
	//如果只有一個結點 且Index是要刪除的值
	if temp.Next == c && temp.Index == index {
		temp.Next = nil
		return c, c, nil
	}

	for {
		//找到最後 而且沒有正確的index
		if temp.Next == c && temp.Index != index {
			err = errors.New("Index not found")
			return
		}
		if c.Index == index {
			c = c.Next
			helper.Next = c
			return c, c, nil
		}
		if temp.Index == index {
			fmt.Println("找到", index)
			helper.Next = temp.Next
			return c, temp, nil
		}

		temp = temp.Next
		helper = helper.Next
	}
}

func (c *CirLnLs) NewHelper() (helper *CirLnLs) {
	helper = c
	for {
		if helper.Next == c {
			return
		} else {
			helper = helper.Next
		}
	}
}

func (c *CirLnLs) ShowCirLnLs() {
	temp := c
	if c.Next == nil {
		fmt.Println("環形鏈表為空")
		return
	}
	if temp == temp.Next {
		fmt.Println("Index: ", temp.Index, "Name: ", temp.Name, "Next:", temp.Next)
		return
	}
	for {
		if temp.Next != c {
			fmt.Println("Index: ", temp.Index, "Name: ", temp.Name, "Next:", temp.Next)
			temp = temp.Next
		} else {
			fmt.Println("Index: ", temp.Index, "Name: ", temp.Name, "Next:", temp.Next)
			return
		}
	}
}
func main() {
	//創建一隻貓
	c1 := NewCirLnLs(1, "湯姆貓")
	c2 := NewCirLnLs(2, "傑立鼠")
	c3 := NewCirLnLs(3, "米老鼠")
	c4 := NewCirLnLs(4, "飛天鼠")
	c5 := NewCirLnLs(5, "穿山鼠")
	c1.InsertNode(c2)
	c1.InsertNode(c3)
	c1.InsertNode(c4)
	c1.InsertNode(c5)
	c1.ShowCirLnLs()
	fmt.Println("-------刪除任務----------")
	c1, _, err := c1.DeleteNode(1)
	c1, _, err = c1.DeleteNode(2)
	c1, _, err = c1.DeleteNode(3)
	c1, _, err = c1.DeleteNode(4)
	c1, _, err = c1.DeleteNode(5)
	// fmt.Println("haha:", c1, c4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("-------顯示任務----------")
		c1.ShowCirLnLs()
	}

}
