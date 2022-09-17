package main

import "fmt"

func main() {

	var multiArray = [2][2]int{
		{14, 27}, 
		{30, 12},
	}

	fmt.Println(multiArray)
	fmt.Println(multiArray[0])		// [14, 27]
	fmt.Println(multiArray[1][0])	// 30
	fmt.Println(multiArray[1][1])	// 12
	fmt.Println(multiArray[0][0])	// 14

//*==========================================================
	fmt.Println("=============================")

	var multiArray1 = [4][3]int{
		{14, 27, 35}, 
		{30, 12, 63},
		{78, 45, 39}, 
		{31, 42, 63},
	}

	fmt.Println(multiArray)
	fmt.Println(multiArray1[0])		// [14, 27, 35]
	fmt.Println(multiArray1[0][1])	// 27
	fmt.Println(multiArray1[2][1])	// 45
	fmt.Println(multiArray1[3][2])	// 63

}