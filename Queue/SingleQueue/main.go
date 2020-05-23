package main

import (
	"errors"
	"fmt"
	"os"
)

//建立一個隊列結構體
type SingleQueue struct {
	MaxSize int
	Front   int
	Rear    int
	Array   []interface{}
}

func NewQueue(l int) *SingleQueue {
	q := new(SingleQueue)
	q.MaxSize = l
	q.Rear = -1
	q.Front = -1
	q.Array = make([]interface{}, l)
	return q
}

//添加數據到隊列
func (q *SingleQueue) AddQueue(d interface{}) (err error) {
	//先判斷隊列是否已滿
	if q.Rear == q.MaxSize-1 { //rear是隊列尾部 (含最後的元素))
		return errors.New("Queue is full.")
	}

	q.Rear++ //rear 往後移
	q.Array[q.Rear] = d
	return nil
}

//顯示隊列 找到隊首 遍歷到隊尾
func (q *SingleQueue) ShowQueue() {
	fmt.Println("隊列當前的情況是:")
	//front不包含隊首的元素
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("array[%d]=%v\n", i, q.Array[i])
	}
}
func (q *SingleQueue) GetQueue() (val interface{}, err error) {
	//先判斷隊列是否為空
	if q.Rear == q.Front {
		return nil, errors.New("Queue Empty")
	}
	q.Front++
	val = q.Array[q.Front]
	return
}
func main() {
	q := NewQueue(5)
	var k string
	for {
		fmt.Println(`
      1.添加數據
      2.獲取數據
      3.顯示隊列
      4.quit`)
		fmt.Scan(&k)
		switch k {
		case "1":
			var v string
			fmt.Println("輸入要加入的值")
			fmt.Scan(&v)
			err := q.AddQueue(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("新增成功")
		case "2":
			val, err := q.GetQueue()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("取出: ", val)
		case "3":
			q.ShowQueue()
		case "4":
			os.Exit(1)
		default:
			fmt.Println("Wrong")
		}
	}
}
