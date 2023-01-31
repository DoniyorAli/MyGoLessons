package main

import "fmt"

func main() {

	var num1, num2, num3 float32
	var bigNum, smalNum float32
	fmt.Println("Enter 3 random numbers: ")
	fmt.Scanln(&num1, &num2, &num3) // 11, 15, 13

	bigNum = num1        // --> 11
	if bigNum < num2{    // --> 11 < 15
		bigNum = num2    // --> 15
	}
	if bigNum < num3{    // --> 15 < 13 
		bigNum = num3  	 // error
	}
	fmt.Println("Big number: ", bigNum) // --> 15

	smalNum = num1       // --> 11
	if smalNum > num2{   // --> 11 > 15
		smalNum = num2   // error
	}
	if smalNum > num3{   // --> 11 > 13
		smalNum = num3   // error
	}
	fmt.Println("Small number: ", smalNum)
}
