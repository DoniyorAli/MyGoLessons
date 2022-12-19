package main

import "fmt"

func main() {

	number := 0
	fmt.Println("Even or odd programm, enter the number: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println(number, "Even number")
	} else {
		fmt.Println(number, "An odd number")
	}

}
