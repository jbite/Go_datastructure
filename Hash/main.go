package main

import (
	"fmt"
	"os"
)

//使用hash來創建一個一存放員工資料的數據結構

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//方法待定

//定義EmpLink
//這裡的Emplinl不帶表頭
//第一個節點就存放雇員
type EmpLink struct {
	Head *Emp
}

//添加員工的方法 保證添加時 編號從小到大
func (empl *EmpLink) Insert(emp *Emp) {
	cur := empl.Head   //這是一個輔助指針
	var pre *Emp = nil //這是一個輔助指針 pre始終在cur前面
	//如果當前的EmpLink就是一個空鏈表
	if cur == nil {
		empl.Head = emp
		return
	}
	//如果不是空鏈表 給emp找到對的位置並插入
	//思路是 讓cur和emp比較 讓pre保持在cur前面
	for {
		if cur != nil {
			if cur.Id > emp.Id && empl.Head == cur {
				emp.Next = cur
				empl.Head = emp
				return
			}
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	//退出時 決定是否在鏈表最後
	pre.Next = emp
	emp.Next = cur
}

//查找
func (empl *EmpLink) FindByID(id int) *Emp {
	if empl.Head == nil {
		return nil
	}
	temp := empl.Head
	for {
		if temp.Id == id && temp != nil {
			return temp
		} else if temp == nil {
			return nil
		} else {
			temp = temp.Next
		}
	}

}

//顯示linkLisr
func (empl *EmpLink) ShowLink(no int) {
	if empl.Head == nil {
		fmt.Printf("此鏈表%d為空\n", no+1)
		return
	}
	temp := empl.Head
	fmt.Printf("此鏈表%d內容:", no+1)
	for {
		if temp != nil && temp.Next != nil {
			fmt.Printf("雇員id=%d 名字=%s => ", temp.Id, temp.Name)
			temp = temp.Next
		} else if temp != nil && temp.Next == nil {
			fmt.Printf("雇員id=%d 名字=%s\n", temp.Id, temp.Name)
			break
		} else {
			fmt.Println()
			break
		}
	}
}

type HashTable struct {
	LinkArr [7]EmpLink
}

//給HashTable 編寫inset雇員的方法
func (h *HashTable) Insert(emp *Emp) {
	//使用散列函數確定將該雇員添加到哪個鏈表
	linkNo := h.HashFun(emp.Id)
	//使用對應的鏈表添加
	h.LinkArr[linkNo].Insert(emp)
}

//編寫方法顯示hash表內所有雇員
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

//查找雇員
func (h *HashTable) FindById(id int) (emp *Emp) {
	l := h.HashFun(id)
	return h.LinkArr[l].FindByID(id)
}

//編寫一個散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7 //得到一個值 也就是鏈表的下標
}

func main() {

	h := HashTable{}
	for {
		fmt.Println(`
==============雇員系統菜單================
1. 添加雇員
2. 顯示雇員
3. 查找雇員
4. 退出系統
請輸入選擇: `)

		var key string
		var id int
		var name string
		fmt.Scan(&key)
		switch key {
		case "1":
			fmt.Println("輸入雇員id:")
			fmt.Scan(&id)
			fmt.Println("輸入雇員名字:")
			fmt.Scan(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			h.Insert(emp)
		case "2":
			h.ShowAll()
		case "3":
			fmt.Println("輸入雇員Id")
			fmt.Scan(&id)
			emp := h.FindById(id)
			if emp != nil {
				fmt.Println("Id", id, "是", emp.Name)
			} else {
				fmt.Println("查無此人")
			}
		case "4":
			os.Exit(1)
		default:
			fmt.Println("輸入有誤")
		}
	}
}
