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

	fmt.Println("===>", h1)
	fmt.Println("===>", h1.Name, h1.Surname, h1.Age)

	h2 := Human{Name: "John", Surname: "Doe", Age: 35}
	box := h2.Name + " " + h2.Surname + " " + strconv.Itoa(h2.Age)
	fmt.Println("===>", box)

}
