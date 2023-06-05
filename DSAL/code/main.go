package main

import (
	"fmt"
)

func main() {
	square, cube := powerSeries(3)
	fmt.Println(square, cube)
}

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}
