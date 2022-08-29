package main

import "fmt"

func main() {

	var num1 float32
	var num2 float32
	var arithmetic string
	fmt.Printf("Enter number:")
	fmt.Scan(&num1)
	fmt.Printf("Enter the arifmetiks -; +; *; / : ")
	fmt.Scan(&arithmetic)
	fmt.Printf("Enter number:")
	fmt.Scan(&num2)

	if arithmetic == "+" {
		fmt.Printf("Result: %.1f", num1 + num2)
	}else if arithmetic == "-" {
		fmt.Printf("Result: %.1f", num1 - num2)
	}else if arithmetic == "*" {
		fmt.Printf("Result: %.1f", num1 * num2)
	}else if arithmetic == "/" {
		fmt.Printf("Result: %.1f", num1 / num2)
	}else {
		fmt.Println("Wrong entered!")
	}

}