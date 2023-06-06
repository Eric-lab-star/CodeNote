// Tuples
package main

import "fmt"

func main() {
	fmt.Println(timeTable(2))
}

func timeTable(a int) (int, int) {
	return a, a * 2
}
