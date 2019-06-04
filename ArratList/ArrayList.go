package ArratList

import (
	"fmt"
	"errors"
)

// 接口
type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newval interface{}) error // 插入数据
	Append(newval interface{})             //追加数据
	Clear()                                     // 清空
	Delete(index int) error
	String() string
}

// 数据结构,字符串,整数,
type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

// 开辟内存
func NewArrayList()*ArrayList{
	list:=new(ArrayList) //初始化结构体
	list.dataStore = make([]interface{},0,10)
	list.theSize = 0
	return list
}

// 数组大小
func (list *ArrayList)Size() int{
	return list.theSize //返回数组大小
}
func (list *ArrayList)String() string{
 return fmt.Sprint(list.dataStore)
}
// 追加元素
func (list *ArrayList)Append(newval interface{}){
	list.dataStore = append(list.dataStore,newval)
	list.theSize++
}

// 抓取数据
func (list *ArrayList)Get(index int) (interface{},error){
   	if index<0 || index>=list.theSize{
   		return nil,errors.New("索引越界")
	}

   	return list.dataStore[index],nil
}