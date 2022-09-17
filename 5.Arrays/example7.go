package main

import "fmt"

func main() {

	var multiArray = [4][3]int{
		{1, 7, 3}, 
		{3, 7, 6},
		{8, 5, 9}, 
		{3, 24, 6},
	}

	fmt.Println(multiArray)		// [[14 27 35] [30 12 63] [78 45 39] [31 42 63]]
	fmt.Println(multiArray[0][0] + multiArray[1][1])	// 1 + 7 = 8
	fmt.Println(multiArray[0][1] - multiArray[0][2])	// 7 - 3 = 4
	fmt.Println(multiArray[2][1] * multiArray[1][1])	// 5 * 7 = 35
	fmt.Println(multiArray[3][1] / multiArray[3][2])	// 24 / 6 = 4

}