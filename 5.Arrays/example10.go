package main

import "fmt"

func main() {

	var myArray[3][4] int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			myArray[i][j] = i + j
			fmt.Printf(" %d", myArray[i][j]) //* 0 1 2 3
		}									 //* 1 2 3 4
		fmt.Println()                        //* 2 3 4 5
	}                                        
	fmt.Println(myArray)	//* [[5 9 777 6 4] [3 6 2 4444 0] [12 95 75 455 456]]
}