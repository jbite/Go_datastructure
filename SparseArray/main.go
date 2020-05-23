package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	chessMap [11][11]int
)

type SparseNode struct {
	Row    int
	Column int
	Data   interface{}
}

func NewSparseNode(r, c int, d interface{}) *SparseNode {
	n := new(SparseNode)
	n.Row = r
	n.Column = c
	n.Data = d
	return n
}

type SparseArray struct {
	RowSize int
	ColSize int
	Default int
	Array   []*SparseNode
}

func (s *SparseArray) Save() {
	//(1)遍歷chessMap 如果我們發現有一個元素的值不為0,就創建一個node結構體
	//(2)將其放入到對應的切片中
	//標準的稀疏數組應該還有表示原始二維數組的規模(行和列,默認值)
	for i, v := range chessMap {
		s.RowSize++
		s.ColSize = len(v)
		for j, w := range v {
			if w != 0 {
				cn := NewSparseNode(i, j, w)
				s.Array = append(s.Array, cn)
			}
		}
	}
	//將讀出的結果寫入文件
	var str string
	str = fmt.Sprintf("%d %d %d\n", s.RowSize, s.ColSize, s.Default)
	for _, v := range s.Array {
		s := fmt.Sprintf("%d %d %d\n", v.Row, v.Column, v.Data)
		str += s
	}
	filepath := "sparsedata"
	ioutil.WriteFile(filepath, []byte(str), 0644)

}

func main() {
	//先創建一個原始數組
	// var chessMap [11][11]int
	var sparseArr SparseArray

	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //藍子

	//輸出原始的數組
	for _, v := range chessMap {
		for _, w := range v {
			fmt.Printf("%d\t", w)
		}
		fmt.Println("")
	}
	//將稀疏數組 儲存
	sparseArr.Save()

	var chessMap2 [11][11]int
	//恢復原始的數組
	filepath := "sparsedata"
	file, _ := os.OpenFile(filepath, os.O_RDONLY, 0644)
	defer file.Close()

	reader := bufio.NewReader(file)
	for true {
		str, err := reader.ReadString('\n')
		if err == nil {
			ns := strings.Fields(str)
			v1, _ := strconv.Atoi(ns[0])
			v2, _ := strconv.Atoi(ns[1])
			v3, _ := strconv.Atoi(ns[2])
			if v1 >= len(chessMap2) {
				continue
			} else {
				chessMap2[v1][v2] = v3
			}
		} else if err == io.EOF {
			break
		}

	}
	//輸出chessMap2的數組
	fmt.Println("chessMap2:")
	for _, v := range chessMap2 {
		for _, w := range v {
			fmt.Printf("%d\t", w)
		}
		fmt.Println("")
	}
	fmt.Println()
	//打開這個文件
	//使用稀疏數組恢復
	//

	//遍歷文件的每一行

}
