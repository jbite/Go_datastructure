package main

import (
	"fmt"

	"./ArrayList"
)

func main() {
	//定義接口隊像 必須時線接口所有的方法
	list := ArrayList.NewArrayList()
	var c string
	var i int
	var ele string
	for {
		fmt.Println(`Pleas enter the command:
  1. Show list
  2. Get(index)
  3. Set(index, ele)
  4. Append(ele)
  5. Insert(index, ele)
  6. Clear
  7. Delete`)
		fmt.Scan(&c)
		switch c {
		case "1":
			fmt.Println(list.String())
		case "2":
			fmt.Println("輸入想取得第幾個元素?")
			fmt.Scan(&i)
			e, err := list.Get(i - 1)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(e)
			}

		case "3":
			if list.TheSize == 0 {
				fmt.Println("No element in list")
				break
			}
			fmt.Println(list.String())
			fmt.Println("輸入想設置第幾個元素?")
			fmt.Scan(&i)
			fmt.Println("輸入設置的元素內容?")
			fmt.Scan(&ele)
			// fmt.Println(i)
			err := list.Set(i-1, ele)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(list.String())
			}
		case "4":
			fmt.Println("輸入擴展的元素內容?")
			fmt.Scan(&ele)
			list.Append(ele)
			fmt.Println(list.String())
		case "5":
			fmt.Println(list.String())
			fmt.Println("輸入想插入第幾個元素?")
			fmt.Scan(&i)
			fmt.Println("輸入插入的元素內容?")
			fmt.Scan(&ele)
			err := list.Insert(i-1, ele)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(list.String())
			}
		case "6":
			var confirm string
			fmt.Println("Are you sure?(y/n)")
			fmt.Scan(&confirm)
			if confirm == "y" {
				list.Clear()
			} else {
				fmt.Println("No execute clear")
				break
			}

		case "7":
		default:
			fmt.Println("輸入有誤")
		}
	}

	// fmt.Println(list) //小寫私有 大寫公有
}
