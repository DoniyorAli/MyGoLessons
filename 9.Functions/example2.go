package main

import "fmt"

func main() {

	one, two := split(11)
	fmt.Printf("First: %d\nSecond: %d\n", one, two)
	fmt.Println()
	length, total := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("LengthNum:%d\nTotalNum:%d\n", length, total)

}

//!=========================Functions========================================

func split(num int) (x, y int) {
	x = num + 10
	y = num - x
	return // ---> x, y
}
//*==========================================================================

func add(nums ... int) (length, result int) {
	// result = 0
	for _, value := range nums {
		result += value    // ---> named (result)
	}
	length = len(nums)     // ---> named (length)
	return // ---> length, result
}



