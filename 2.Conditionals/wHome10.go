package main

import "fmt"

func main() {

	var num1, num2, num3, num4, num5 int
	fmt.Println("Enter five numbers:")
	fmt.Scan(&num1, &num2, &num3, &num4, &num5)
	bigNum := num1

	if num1 == num2 && num2 == num3 && num3== num4 && num4 == num5 {
		fmt.Println("All numbers equal")
		
	}
	if bigNum < num2 {
		bigNum = num2
	}
	if bigNum < num3{
		bigNum = num3
	}
	if bigNum < num4 {
		bigNum = num4
	}
	if bigNum < num5 {
		bigNum = num5
	}

	fmt.Println("Big number:", bigNum)
}