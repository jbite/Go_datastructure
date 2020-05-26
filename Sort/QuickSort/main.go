package main

import (
	"fmt"
	"math/rand"
	"time"
)

//left 數組左邊的索引
//right 數組右邊的數
//array表示要排序的數組
func quickSort(left int, right int, array *[8000000]int) {
	l := left
	r := right

	//使用首數當比較點
	pivot := array[(left+right)/2]
	for l < r {

		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if r < l {
			break
		} else {
			if array[l] != array[r] {
				array[l], array[r] = array[r], array[l]
			} else {
				break
			}
		}

		// fmt.Println("pivot:", pivot)
		// fmt.Println("l:", l, "r:", r)
		// fmt.Println("array:", *array)
	}

	if l >= r {
		l--
		r++
	}

	if left < l {
		quickSort(left, l, array)
	}
	if right > r {
		quickSort(r, right, array)
	}
}

func main() {
	var s [8000000]int
	for i := 0; i < 8000000; i++ {
		l := rand.Intn(100000000)
		s[i] = l
	}
	// s := [10]int{13, 5, -100, -100, 46, 9, 6, 31, 20, 45}
	// fmt.Println("origin:", s)
	start := time.Now().Unix()

	quickSort(0, len(s)-1, &s)
	end := time.Now().Unix()
	// for _, v := range s {
	// 	fmt.Println(v)
	// }
	fmt.Println(end - start)
}
