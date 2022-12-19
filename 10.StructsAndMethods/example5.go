//* Circle, Triangle, Square, Rectangle
//* () Perim() float64
//* () Area() float64

package main

import (
	"errors"
	"fmt"
	"math"
)

//*==========================Circle========================
type Circle struct {
	r float64
}
//* Perimetr...
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.r
}
//* Area...
func (c Circle) Area() float64 {
	return math.Pi * float64(c.r * c.r)
}

//*========================Triangle========================
type Triangle struct {
	a, b, c float64
}
//* Perimetr...
func (t Triangle) Perim() float64 {
	return t.a + t.b + t.c
}
//* Area...
func (t Triangle) Area() float64 {
	per := t.Perim() / 2
	area := math.Sqrt(per * (float64(per - t.a) * float64(per - t.b) * float64(per - t.c)))
	return area
}
//* Triangle checking...
func (t Triangle) Check() bool {
	if t.a + t.b <= t.c || t.b + t.c <= t.a || t.a + t.c <= t.b {
		return false
	}
	return true
}

//*========================Square==========================
type Square struct {
	a float32
}
//* Perimetr...
func (s Square) Perim() float64 {
	return float64(4 * s.a)
}
//* Area...
func (s Square) Area() float64 {
	return float64(s.a * s.a)
}

//*========================Rectangle=======================
type Rectangle struct {
	a, b float64
}
//* Perimetr...
func (r Rectangle) Perim() float64 {
	return 2 * (r.a + r.b)
}
//* Area...
func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func main() {

	shapeCircle := Circle {
		r: 7,
	}
	fmt.Printf("Circle\nPerimeter=%f,\nArea=%f\n", shapeCircle.Perim(), shapeCircle.Area())
	fmt.Println("==========================")

	shapeTriangle := Triangle {
		a: 2,
		b: 4,
		c: 5,
	}

	if !shapeTriangle.Check() {
		err := errors.New("This is not Triangle!")
		panic(err)
	}

	fmt.Printf("Triangle\nPerimeter=%f,\nArea=%f\n", shapeTriangle.Perim(), shapeTriangle.Area())
	fmt.Println("==========================")

	shapeSquare := Square {
		a: 5,
	}
	fmt.Printf("Square\nPerimeter=%f,\nArea=%f\n", shapeSquare.Perim(), shapeSquare.Area())
	fmt.Println("==========================")

	shapeRectangle := Rectangle {
		a: 8,
		b: 7,
	}
	fmt.Printf("Rectangle\nPerimeter=%0.f,\nArea=%0.f\n", shapeRectangle.Perim(), shapeRectangle.Area())


}