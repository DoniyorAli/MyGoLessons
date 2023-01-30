package main    // IF; ELSE; ELSE IF

import "fmt"

func main() {

	var A int
	var B int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&A)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&B)

	if A > B {
		fmt.Println("A big number!")
	}else if A == B{
		fmt.Println("Both of them equal!")
	}else {
		fmt.Println("B big number!")
	}
}
