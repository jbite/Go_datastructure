package ArrayList

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()       //刪除
	GetIndex() int //得到索引
}

type Iterable interface {
	Iterator() Iterator //構造初始化街口
}

//構造指針訪問數組
type ArrayListIterator struct {
	list         *ArrayList //數組指針
	currentindex int
}
