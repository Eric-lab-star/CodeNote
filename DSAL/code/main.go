// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape2 struct{}

func (DrawShape2) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape2")
}

// DrawShape struct
type DrawShape struct{}

// DrawShape struct has method draw Shape with float x and y coordinates
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

// IContour interace
type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// DrawContour method resizeByFactor given factor

// DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  IDrawShape
	factor int
}

// DrawContour method drawContour given the coordinates
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// DrawContour method resizeByFactor given factor
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

// main method
func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
