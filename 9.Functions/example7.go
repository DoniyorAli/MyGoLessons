package main

import "fmt"

func main() {

	fmt.Println("Result:", add(11, 22))

	box := name("Ali")
	box()

	name("John")()

}

var add = func(n1, n2 int) (result int) {   //! Maqsad nima variablega olishdan baribir return qivommizku
	result = n1 + n2
	return
}

var name = func(name string) func() {     //! ishlash tartibi qanaqa
	return func() {
		fmt.Println("Hello", name)
	}
}
//? QUESTION

