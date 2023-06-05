package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList = list.List{}
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(int))
	}
}
