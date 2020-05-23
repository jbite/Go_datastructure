package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (interface{}, error)         //抓取第幾個元素
	Set(index int, newval interface{}) error    //修改數據
	Insert(index int, newval interface{}) error //插入數據
	Append(newval interface{})                  //追加數據
	Clear()                                     //清空
	Delete(index int) error                     //刪除數據
	String() string                             //返回字符串
}

type ArrayList struct {
	dataStore []interface{} //數據存儲
	TheSize   int           //數組的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      //初始化結構體
	list.dataStore = make([]interface{}, 0, 10) //開闢空間10個
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize
}

//抓取數據
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("index out of size")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newval interface{}) {

	list.dataStore = append(list.dataStore, newval)
	list.checkisFull()
	list.TheSize++

}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("index out of size")
	}
	list.dataStore[index] = newval //設置
	return nil
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("index out of size")
	}
	list.checkisFull()

	list.dataStore = list.dataStore[:list.TheSize+1]
	// fmt.Printf("%v,%v\n", list.TheSize, len(list.dataStore))
	for i := list.TheSize; i > index; i-- {
		// fmt.Println(i)
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.TheSize++
	return nil
}

func (list *ArrayList) checkisFull() {
	if list.TheSize == cap(list.dataStore) { //判斷內存使用
		newdataStore := make([]interface{}, 0, 2*list.TheSize)
		// copy(newdataStore, list.dataStore)
		newdataStore = append(newdataStore, list.dataStore...) //拷貝
		list.dataStore = newdataStore
	}
}
func (list *ArrayList) Delete(index int) error { //刪除
	if index < 0 || index >= list.TheSize {
		return errors.New("index out of size")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.TheSize--

	return nil
}
