package main

import "fmt"

func main() {

	var num1, num2, num3, total, multip, arithmetik, bigNum, smalNum float32
	fmt.Println("Enter the numbers: ")
	fmt.Scanln(&num1, &num2, &num3)

	total = num1 + num2 + num3
	multip = num1 * num2 * num3
	arithmetik = total / 3

//  Big number
	smalNum = num1
	if smalNum > num2 {
		smalNum = num2
	}
	if smalNum > num3{
		smalNum = num3
	}
// Small number
	bigNum = num1
	if bigNum < num2{
		bigNum = num2
	}
	if bigNum < num3{
		bigNum = num3
	}

	fmt.Println("Total:", total, "\nMultip:", multip, "\nArithmetik:", arithmetik)
	fmt.Println("Big number:", bigNum, "\nSmall number:", smalNum)

}