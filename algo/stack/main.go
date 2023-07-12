package main

import "fmt"

type stack []int

func (s *stack) Pop() int {
	ele := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ele
}

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func main() {
	s := &stack{}
	s.Push(1)
	s.Push(2)
	s.Pop()
	fmt.Println(s)
}
