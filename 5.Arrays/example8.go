
package main

import "fmt"

func main() {

	var multiArray = [4][3]int{{1, 7, 3}, {3, 7, 6},{8, 5, 9}, {3, 24, 6},}

	fmt.Println(multiArray)		//* [[1 7 3] [3 7 6] [8 5 9] [3 24 6]]
	multiArray[0][1] = 2
	multiArray[2][1] = 9
	multiArray[2][2] = 10
	multiArray[3][0] = 10
	multiArray[3][1] = 11
	multiArray[3][2] = 12
	fmt.Println(multiArray)		//* [[1 2 3] [3 7 6] [8 9 10] [10 11 12]]

}