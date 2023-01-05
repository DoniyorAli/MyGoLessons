package main

import "fmt"

func main() {

	number1 := 0
	number2 := 0
	number3 := 0

	fmt.Println("Enter first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&number2)
	fmt.Println("Enter third number: ")
	fmt.Scanln(&number3)

	if number2 < number1 && number1 > number3 {
		fmt.Printf("Big number: %d\n", number1)
	}
	if number1 < number2 && number2 > number3 {
		fmt.Printf("Big number: %d\n", number2)
	}
	if number1 < number3 && number3 > number2 {
		fmt.Printf("Big number: %d\n", number3)
	}
	if number1 == number2 && number2 == number3 {
		fmt.Println("Equal")
	} 

	// bigNum := number1
	// if bigNum < number2 {
	// 	bigNum = number2
	// }
	// if bigNum < number3 {
	// 	bigNum = number3
	// }
	// if number1 == number2 && number2 == number3{
	// 	fmt.Println("All numbers are equal!")
	// }
	// fmt.Println("The largset number:", bigNum)

}
