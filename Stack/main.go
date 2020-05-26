package Stack

import (
	"errors"
	"fmt"
)

//使用數組來模擬棧的使用

type Stack struct {
	maxTop int
	Top    int
	Data   [20]int
}

func NewStack() *Stack {
	s := new(Stack)
	s.maxTop = 20
	s.Top = -1
	return s
}

func (s *Stack) Push(val int) (err error) {
	//先判斷stack是否滿了
	if s.Top == s.maxTop-1 {
		fmt.Println("Stack is full")
		return errors.New("Stack is full")
	}
	s.Top++
	s.Data[s.Top] = val
	return
}

//出棧
func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("Stack is empty")
		return 0, errors.New("Stack is empty")
	}

	val = s.Data[s.Top]
	s.Top--
	return
}

//遍歷棧
func (s *Stack) ShowStack() (err error) {
	if s.Top == -1 {
		fmt.Println("stack is empty")
		return errors.New("stack is empty")
	}
	fmt.Println("Stack:")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("Data[%d]=%d\n", i, s.Data[i])
	}
	return
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	v, _ := s.Pop()
	fmt.Println("取出", v)
	s.ShowStack()
}
