package main

import (
	"fmt"
	"time"
)

func SelectSort(s []int) []int {
	//比較頭到尾巴前一個 最後一個不用再比
	for i := 0; i < len(s)-1; i++ {
		max := s[i]
		maxIndex := 0
		//比較i+1以後到尾端的數字
		for j := i + 1; j < len(s); j++ {
			if max < s[j] {
				max = s[j]
				maxIndex = j
			}
		}
		if maxIndex != 0 {
			s[i], s[maxIndex] = s[maxIndex], s[i]
		}
	}
	return s
}
func main() {
	s := []int{13, 5, 444, 77, 46, 93, 2, 9, 11, 6}
	starttime := time.Now()
	ns := SelectSort(s)
	endtime := time.Now()

	fmt.Printf("耗時:%v\n", endtime.Sub(starttime))
	fmt.Println(ns)
}
