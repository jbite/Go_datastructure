package main

import (
	"errors"
	"fmt"
)

//使用帶head頭的單向鏈表實現水滸英雄排行榜管理

type LinkList struct {
	Index    int
	Name     string
	NickName string
	Next     *LinkList
}

func HeadLinkList() *LinkList {
	list := new(LinkList)
	return list
}
func NewLinkList(index int, name string, nickname string) *LinkList {
	list := new(LinkList)
	list.Index = index
	list.Name = name
	list.NickName = nickname
	return list
}

//Push 給鏈表插入一個結點
//編寫第一種插入方法 在單鏈表的最後加入
//1. 先找到該鏈表的最後這個結點
//2. 創建輔助結點[跑腿的]
//3.如果Next為空 就將新結點加入
func (h *LinkList) Push(newnode *LinkList) {
	temp := h
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next //讓temp不斷的指向下一個節點
	}
	temp.Next = newnode //將結點指向新結點
}

//SortPush 編寫第2種插入方法 根據no的編號從小到大插入
//讓temp在適當的位置先插入結點，因為跳過後 temp無法再回頭
func (h *LinkList) SortPush(newnode *LinkList) (err error) {
	temp := h
	if temp.Next == nil { //已經到鏈表的最後了
		temp.Next = newnode
		return
	} else if temp.Next.Index == newnode.Index {
		err = errors.New("Index ruplicate")
		fmt.Println(newnode.Name, err)
		return
	} else if newnode.Index < temp.Next.Index {
		//就應該插入下一跳的前面
		newnode.Next = temp.Next
		temp.Next = newnode
	} else if newnode.Index > temp.Next.Index {
		//當newnode index比現在結點的indez大就遞歸
		err = temp.Next.SortPush(newnode)
		return
	}
	return
}

//DeleteNode 刪除結點
func (h *LinkList) DeleteNode(index int) (err error) {
	temp := h
	for {
		if temp.Next == nil {
			err = errors.New("結點不存在")
			fmt.Println(index, err)
			return
		} else if temp.Next.Index != index {
			temp = temp.Next
		} else if temp.Next.Index == index {
			temp.Next = temp.Next.Next
			fmt.Println(temp.Next.Name, "已經刪除")
			return
		}
	}
}

func (head *LinkList) ShowLinkList() {
	temp := head
	for {
		if temp.Next != nil {
			fmt.Printf("Index:%3v\t, Name:%-5s\t, Nickname: %-5s\n",
				temp.Next.Index,
				temp.Next.Name,
				temp.Next.NickName)
			temp = temp.Next //將結點指向下一個節點
		} else {
			break
		}
	}
	return
}
func main() {
	head := HeadLinkList()
	hero1 := NewLinkList(6, "宋江", "及時雨")
	hero2 := NewLinkList(3, "武松", "行者")
	hero3 := NewLinkList(5, "盧俊義", "玉麒麟")
	hero4 := NewLinkList(7, "林沖", "豹子頭")
	hero5 := NewLinkList(101, "吳用", "智多星")
	hero6 := NewLinkList(1, "公孫勝", "入雲龍")
	head.SortPush(hero1)
	head.SortPush(hero2)
	head.SortPush(hero3)
	head.SortPush(hero4)
	head.SortPush(hero5)
	head.SortPush(hero6)
	head.ShowLinkList()
	head.DeleteNode(5)
	head.DeleteNode(5)
	// fmt.Printf("%#v\n", head)
	head.ShowLinkList()
}
