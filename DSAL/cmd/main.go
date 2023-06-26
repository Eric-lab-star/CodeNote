package main

import "fmt"

/*
Tuples are finit orderd sequences of objects. They can contain a mixture of other data types and are used to group related data into a data structure. In a relational database, a tuple is a row of a table. Tuples have a fixed size compared to lists, and are also faster.
*/

func h(x, y int) int{
	return x*y
}

func g(l,m int) (x,y int){
	x= 2*l
	y=4*m
	return
}
func main(){
	fmt.Println(h(g(2,5)))
}
