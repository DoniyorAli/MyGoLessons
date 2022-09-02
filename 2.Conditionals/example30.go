package main

import "fmt"

func main() {

	var first int
	var actions string
	var second int

	fmt.Print("Enter the first number:")
	fmt.Scanf("%v", &first)
	fmt.Print("Enter the actions +, -, /, * :")
	fmt.Scanf("%v", &actions)
	fmt.Print("Enter the second number:")
	fmt.Scanf("%v", &second)

	if actions == "+" {
		fmt.Println("Total:", first + second)
	}
	if actions == "-" {
		fmt.Println("Total:", first - second)
	}
	if actions == "*" {
		fmt.Println("Total:", first * second)
	}
	if actions == "/" {
		fmt.Println("Total:", first / second)
	}

}