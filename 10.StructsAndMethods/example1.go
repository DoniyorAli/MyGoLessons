package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	
	h1 := Human{Name: "Ali", Surname: "Khanov", Age: 27}
	box := h1.Name + " " + h1.Surname + " " + strconv.Itoa(h1.Age)
	fmt.Println(box)

	h2 := Human{Name: "John", Surname: "Bon", Age: 29}
	fmt.Println(h2)
//*============================================================================
	fmt.Println()
	// var h3 = new(Human)	//* Empty Human 
	// fmt.Println(h3)

	var h3 = new(Human)
	h3.Name = "Ali"
	h3.Surname = "Khanov"
	h3.Age = 27
	fmt.Println(h3)
	fmt.Println()

	h4 := NewHuman()
	fmt.Println(h4)
	fmt.Println()

	h5 := NewHumanParams("Daniel", "Saber", 31)
	fmt.Println(h5)
}

//*========================Functions===========================
//! CONSTRUCTOR
func NewHuman() *Human {
	n := new(Human)
	return n
}

//! PARAMETR CONSTRUCTOR
func NewHumanParams(name, surname string, age int) *Human {
	n := new(Human)
	fmt.Println("ParamConstructor")
	n.Name = name
	n.Surname = surname
	n.Age = age
	return n
}
