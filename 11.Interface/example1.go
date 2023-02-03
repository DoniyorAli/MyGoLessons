//* Circle, Triangle, Square, Rectangle
//* () Perim() float64
//* () Area() float64

package main

import (
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

//*=======================INTERFACE========================
type Shape interface {
	Perim() float64
	Area() float64
}

func main() {

	Calculate(Circle{
		r: 7,
	})
	
	Calculate(Triangle{
		a: 3,
		b: 5,
		c: 7,
	})

	Calculate(Square{
		a: 4,
	})

	Calculate(Rectangle{
		a: 4,
		b: 5,
	})

	// if !shapeTriangle.Check() {
	// 	err := errors.New("This is not Triangle!")
	// 	panic(err)
	// }

}

func Calculate(obj Shape) {
	switch obj.(type) {
	case Circle:
		fmt.Println("Circle ==>")
	case Triangle:
		fmt.Println("Circle ==>")
	case Square:
		fmt.Println("Circle ==>")
	case Rectangle:
		fmt.Println("Circle ==>")
	default:
		fmt.Println("=====>Unknown<=====")
	}
	fmt.Printf("Perimeter=%1.f,\nArea=%1.f\n", obj.Perim(), obj.Area())
	fmt.Println("==========================")
}