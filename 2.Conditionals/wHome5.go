package main

import "fmt"

func main() {

	var number1 int 
	var number2 int

	fmt.Printf("Enter first number: ")
	fmt.Scan(&number1)
	fmt.Printf("Enter second number: ")
	fmt.Scan(&number2)

	if number1 > number2{
		fmt.Println("Bg number:", number1)
	}else if number1 == number2{
		fmt.Println("Both of them equal")
	}else {
		fmt.Println("Big number: ", number2)
	}
}