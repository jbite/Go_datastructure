package main

import (
	"errors"
	"fmt"
	"strconv"

	ls "../CircleLinkList"
)

//AddPerson 將玩家加入遊戲 並且返回頭一個玩家
func AddPerson(n int) (head *ls.CirLnLs) {
	name := "a"
	head = ls.NewCirLnLs(1, name+strconv.Itoa(1))
	fmt.Println(head.Name, "加入遊戲了")
	for i := 1; i < n; i++ {
		//創建玩家 並將玩家加入環形鏈表
		d := ls.NewCirLnLs(i+1, name+strconv.Itoa(i+1))
		head.InsertNode(d)
		fmt.Println(d.Name, "加入遊戲了")
	}
	return
}

//CountAndDelete 從選擇的玩家開始 依照約定的跳數報數
//當報數==跳數時，將玩家剔除
//返回剔除的玩家及err
func CountAndDelete(head *ls.CirLnLs, n int) (delNode *ls.CirLnLs, err error) {
	fmt.Println("開始報數:")

	if head == head.Next {
		err = errors.New("剩下一個人在隊裡，遊戲結束了")
		return
	}
	temp := head
	helper := head.NewHelper()
	for i := 0; i < n-1; i++ {
		if temp != temp.Next {
			fmt.Println(temp.Name, "報數:", i+1)
			temp = temp.Next
			helper = helper.Next
		}
	}
	if temp != temp.Next {
		fmt.Println(temp.Name, "報數:", n)
		fmt.Println(temp.Name, "阿~~我死了")
		delNode = temp
		helper.Next = temp.Next

	}
	return
}
func gameStart(total int, hop int) (array []*ls.CirLnLs, last *ls.CirLnLs) {
	if total == 1 {
		fmt.Println("人數不足")
		return
	}
	p := AddPerson(total)
	fmt.Println("加入遊戲的有:")
	p.ShowCirLnLs()
	fmt.Println("每一輪數", hop, "次，數中的人出列")
	last = p
	for {
		del, err := CountAndDelete(last, hop)
		if del != nil {
			array = append(array, del)
		}

		if err != nil {
			fmt.Println("遊戲結數!!!")
			break
		}
		fmt.Println("下一輪由", del.Next.Name, "開始數")
		last = del.Next
	}
	fmt.Println("出局順序為:")
	for i, v := range array {
		fmt.Println(i+1, v.Name)
	}
	fmt.Println("獲勝的人是:", last.Name)
	return
}
func main() {
	gameStart(10, 3)

	// fmt.Println(d.Next)
}
