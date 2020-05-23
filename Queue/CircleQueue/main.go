package main

import (
	"errors"
	"fmt"
	"os"
)

//思路
//隊列滿的時候 就是(tail +1)%maxSize == head
//tail == head 時，隊列為空
//初始時 tail == head == 0
//如何統計 (tail+maxSize-head)%maxSize

type CircleQueue struct {
	MaxSize int
	Head    int
	Tail    int
	Array   []interface{}
}

func NewCircleQueue(l int) *CircleQueue {
	c := new(CircleQueue)
	c.MaxSize = l
	c.Head = 0
	c.Tail = 0
	c.Array = make([]interface{}, l)
	return c
}

//推值入隊列
func (c *CircleQueue) Push(d interface{}) (err error) {
	//判斷隊列是否已滿
	if c.IsFull() {
		err = errors.New("隊列已滿")
		return
	}
	c.Array[c.Tail] = d
	c.Tail++
	if c.Tail == c.MaxSize {
		c.Tail = 0
	}
	return
}

//彈值出隊列
func (c *CircleQueue) Pop() (d interface{}, err error) {
	//判斷隊列是否為空
	if c.IsEmpty() {
		err = errors.New("隊列為空")
		return d, err
	}

	//把值取出隊列
	d = c.Array[c.Head]
	c.Head++
	if c.Head == c.MaxSize {
		c.Head = 0
	}

	return
}

//判斷隊列是否已滿
func (c *CircleQueue) IsFull() bool {
	return (c.Tail+1)%c.MaxSize == c.Head
}

//判斷隊列是否空
func (c *CircleQueue) IsEmpty() bool {
	return c.Tail == c.Head
}

//取出環形隊列元素個數
func (c *CircleQueue) Size() int {
	return (c.Tail + c.MaxSize - c.Head) % c.MaxSize
}

func (c *CircleQueue) ShowCircleQueue() {
	//判斷隊列是否為空
	if c.IsEmpty() {
		fmt.Println("Array is empty")
		return
	}
	for i := c.Head; i < c.Head+c.Size(); i++ {
		//新增一個輔助變數
		a := i % c.MaxSize
		fmt.Printf("Array[%d]: %v\n", i-c.Head+1, c.Array[a])
	}
}

func main() {
	q := NewCircleQueue(6)
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
			err := q.Push(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("新增成功")
		case "2":
			val, err := q.Pop()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("取出: ", val)
		case "3":
			q.ShowCircleQueue()
			fmt.Println(q)
		case "4":
			os.Exit(1)
		default:
			fmt.Println("Wrong")
		}
	}
}
