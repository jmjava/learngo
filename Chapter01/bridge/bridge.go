//main package has examples shown
// in Go Data Structures and algorithms book
package bridge

// importing fmt package
import (
	"fmt"
)

//IDrawShape interface
type IDrawShape interface {
	DrawShape(x [5]float32, y [5]float32)
}

//DrawShape struct
type DrawShape struct{}

// DrawShape struct has  method draw Shape with float x and y coordinates
func (drawShape DrawShape) DrawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

//IContour interace
type IContour interface {
	DrawContour(x [5]float32, y [5]float32)
	ResizeByFactor(factor int)
}

//DrawContour struct
type DrawContour struct {
	X      [5]float32
	Y      [5]float32
	Shape  DrawShape
	Factor int
}

//DrawContour method drawContour given the coordinates
func (contour DrawContour) DrawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.Shape.DrawShape(contour.X, contour.Y)
}

//DrawContour method resizeByFactor given factor
func (contour DrawContour) ResizeByFactor(factor int) {
	contour.Factor = factor
}

