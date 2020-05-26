package main

import (
	"errors"
	"fmt"
)

//使用帶head頭的單向鏈表實現水滸英雄排行榜管理
//聲明雙向列表結構
type DuLinkList struct {
	Index    int
	Name     string
	NickName string
	Prev     *DuLinkList //指向前一個節點
	Next     *DuLinkList //指向下一個節點
}

func HeadDuLinkList() *DuLinkList {
	list := new(DuLinkList)
	return list
}
func NewDuLinkList(index int, name string, nickname string) *DuLinkList {
	list := new(DuLinkList)
	list.Index = index
	list.Name = name
	list.NickName = nickname
	return list
}

func (h *DuLinkList) Push(newnode *DuLinkList) {
	temp := h
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next //讓temp不斷的指向下一個節點
	}
	newnode.Prev = temp
	temp.Next = newnode //將結點指向新結點
}

func (h *DuLinkList) DuSortPush(newnode *DuLinkList) (err error) {
	temp := h
	if temp.Next == nil { //已經到鏈表的最後了
		newnode.Prev = temp
		temp.Next = newnode
		return
	} else if temp.Next.Index == newnode.Index {
		err = errors.New("Index ruplicate")
		fmt.Println(newnode.Name, err)
		return
	} else if newnode.Index < temp.Next.Index {
		//就應該插入下一跳的前面
		newnode.Next = temp.Next
		newnode.Prev = temp
		temp.Next.Prev = newnode
		temp.Next = newnode
	} else if newnode.Index > temp.Next.Index {
		//當newnode index比現在結點的indez大就遞歸
		err = temp.Next.DuSortPush(newnode)
		return
	}
	return
}

//DeleteNode 刪除結點
func (h *DuLinkList) DuDeleteNode(index int) (err error) {
	temp := h
	for {
		if temp.Next == nil && temp.Index != index {
			err = errors.New("結點不存在")
			fmt.Println(index, err)
			return
		} else if temp.Index != index {
			temp = temp.Next
		} else if temp.Index == index {

			if temp.Next != nil {
				temp.Next.Prev = temp.Prev
				temp.Prev.Next = temp.Next
			} else {
				fmt.Println(temp.Prev)
				temp.Prev.Next = HeadDuLinkList().Next
			}

			fmt.Println(temp.Name, "已經刪除")
			return
		}
	}
}

func (head *DuLinkList) ShowDuLinkList() {
	temp := head
	for {
		if temp.Next != nil {
			fmt.Printf("Index:%3v\t, Name:%-5s\t, Nickname: %-5s\n",
				temp.Next.Index,
				temp.Next.Name,
				temp.Next.NickName,
			)
			temp = temp.Next //將結點指向下一個節點
		} else {
			break
		}
	}
	return
}

func (head *DuLinkList) ReverseShowDuLinkList() {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	for {
		if temp != head {
			fmt.Printf("Index:%3v\t, Name:%-5s\t, Nickname: %-5s\n",
				temp.Index,
				temp.Name,
				temp.NickName,
			)
			temp = temp.Prev //將結點指向上一個節點
		} else {
			return
		}
	}
}
func main() {
	head := HeadDuLinkList()
	hero1 := NewDuLinkList(6, "宋江", "及時雨")
	hero2 := NewDuLinkList(103, "武松", "行者")
	hero3 := NewDuLinkList(5, "盧俊義", "玉麒麟")
	hero4 := NewDuLinkList(7, "林沖", "豹子頭")
	hero5 := NewDuLinkList(1, "吳用", "智多星")
	hero6 := NewDuLinkList(2, "公孫勝", "入雲龍")
	head.DuSortPush(hero1)
	head.DuSortPush(hero2)
	head.DuSortPush(hero3)
	head.DuSortPush(hero4)
	head.DuSortPush(hero5)
	head.DuSortPush(hero6)
	head.ShowDuLinkList()
	head.DuDeleteNode(103)
	head.ShowDuLinkList()

}
