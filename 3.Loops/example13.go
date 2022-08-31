package main     // FIBONACCI SERIES

import "fmt"

func main() {
	
	var number int
	fmt.Print("Enter the number:")
	fmt.Scanln(&number)
	x := 1
	y := 0
	z := 0
	for i := 0; i < number; i++ {
		z = x + y
		x = y
		y = z
		fmt.Print(x, " ")
	}
}