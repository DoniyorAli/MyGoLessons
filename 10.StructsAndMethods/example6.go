package main

import (
	"errors"
	"fmt"
	"math"
)

//*=====================================================

type Circle struct {
	r float64
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.r
}
func (c Circle) Area() float64 {
	return math.Pi * float64(c.r * c.r)
}

//*=====================================================

type Triangle struct {
	a, b, c float64
}
func (t Triangle) Perim() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) Area() float64 {
	per := t.Perim() / 2
	area := math.Sqrt(per * float64(per - t.a) * float64(per - t.b) * float64(per - t.c))
	return area
}

func (t Triangle) Check() error {
	var err error
	if t.a + t.b <= t.c || t.b + t.c <= t.a || t.a + t.c <= t.b{
		err = errors.New("This is not Triangle")	//* return errors.New("This is not triangle") ikkinchi yo'li
		return err
	}
	return err //* === nil
}

//*=====================================================

type Square struct {
	a float64
}
func (s Square) Perim() float64 {
	return 4 * s.a
}
func (s Square) Area() float64 {
	return s.a * s.a
}

//*=====================================================

type Rectangle struct {
	a, b float64
}
func (r Rectangle) Perim() float64 {
	return 2 * (r.a + r.b)
}
func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func main() {

	figurCircle := Circle {
		r: 7,
	}
	fmt.Printf("Circle\nPerimeter=%f,\nArea=%f\n", figurCircle.Perim(), figurCircle.Area())
	fmt.Println("==========================")

	figurTriangle := Triangle {
		a: 2,
		b: 3,
		c: 6, 
	}

	err := figurTriangle.Check()

	if err != nil{
		panic(err)
	}

	fmt.Printf("Circle\nPerimeter=%f,\nArea=%f\n", figurTriangle.Perim(), figurTriangle.Area())
	fmt.Println("==========================")

	figurSquare := Square {
		a: 7, 
	}
	fmt.Printf("Circle\nPerimeter=%f,\nArea=%f\n", figurSquare.Perim(), figurSquare.Area())
	fmt.Println("==========================")

	figurRectangle := Rectangle {
		a: 7,
		b: 5, 
	}
	fmt.Printf("Circle\nPerimeter=%f,\nArea=%f\n", figurRectangle.Perim(), figurRectangle.Area())
	fmt.Println("==========================")

}