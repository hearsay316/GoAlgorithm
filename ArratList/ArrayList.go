package ArratList

// 接口
type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newval interface{}) error // 插入数据
	Append(newval interface{}) error            //追加数据
	Clear()                                     // 清空
	Delete(index int) error
	String() string
}

// 数据结构,字符串,整数,
type ArrayList struct {
	dataStore []interface{}
	theSize   int
}
