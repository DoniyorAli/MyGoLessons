package main

import "fmt"

func main() {

	multiArray := [...][5]int{
		{5, 9, 7, 6, 4},
		{3, 6, 2, 4},
		{12, 95, 75, 123, 456},
	}

	fmt.Println(multiArray)		//* [[5 9 7 6 4] [3 6 2 4 0] [12 95 75 123 456]]
	multiArray[2][3] = 455
	fmt.Println(multiArray)		//* [[5 9 7 6 4] [3 6 2 4 0] [12 95 75 455 456]]
	multiArray[0][2] = 777
	fmt.Println(multiArray)		//* [[5 9 777 6 4] [3 6 2 4 0] [12 95 75 455 456]]
	multiArray[1][3] = 4444
	fmt.Println(multiArray)		//* [[5 9 777 6 4] [3 6 2 4444 0] [12 95 75 455 456]]

}