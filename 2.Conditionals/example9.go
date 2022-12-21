package main

import "fmt"

func main() {
//  A program that determines whether or not is divisible without a remainer:

	bigNum, smalNum := 0, 0
	fmt.Printf("Enter the big number: ")
	fmt.Scan(&bigNum)
	fmt.Printf("Enter the small number: ")
	fmt.Scan(&smalNum)

	if bigNum % smalNum == 0{
		fmt.Println("Good this number is devided!")
	}else {
		fmt.Println("Are not devided!")
	}
}
