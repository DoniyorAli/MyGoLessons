package main

import "fmt"

func main() {

	cities := [...]string{"Istanbul", "New York", "Berlin", "Moscow", "Washington"}

	for i, v := range cities {
		fmt.Println("Index:", i, " ", "Value:", v)
	}
	fmt.Println()
	for _, v := range cities {
		fmt.Println("Value:", v)
	}

}