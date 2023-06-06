package main

import (
	"container/list"
	"fmt"
)

func main() {
	intList := list.New()
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)
	intList.PushBack(4)

	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(int))
	}
}
