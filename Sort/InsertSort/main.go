package main

import "fmt"

// "../ArrayList/Array"
func InsertSort(arr [5]int) {
	//完成第一次 給第二個元素找到合適的位置插入
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //數據後移
			insertIndex--                         //index再往前找
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println(arr)
	}
}

func main() {
	s := [5]int{13, 5, 444, 77, 46}

	InsertSort(s)
}
