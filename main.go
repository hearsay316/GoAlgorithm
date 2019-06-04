package main

import (
	"./ArratList"
	"fmt"
)

func main() {
	list := ArratList.NewArrayList()
	list.Append(1)
	list.Append("xwdw")
	fmt.Print(list)
}
