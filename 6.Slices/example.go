package main

import "fmt"

func main() {

	myArray := [...]int{11, 22, 33}
	mySlice := myArray[:]
	fmt.Println("SLICE:", mySlice)	// SLICE: [11 22 33]
	mySlice[0] = 99
	fmt.Println("SlICE:", mySlice)	// SlICE: [99 22 33]
	mySlice = append(mySlice, 456)
	fmt.Println("SlICE:", mySlice)  // SlICE: [99 22 33 456]
//================================================================
	fmt.Println()

	myColors := [...]string{"Red", "Blue", "Yellow", "Black", "White"}
	myColorsSlice1 := myColors[:]
	myColorsSlice2 := myColors[1:len(myColors)-1]
	fmt.Println("Color Slices1:", myColorsSlice1)	// Color Slices1: [Red Blue Yellow Black White]
	fmt.Println("Color Slices2:", myColorsSlice2)	// Color Slices2: [Blue Yellow Black]
	myColorsSlice1 = append(myColorsSlice1, "Hello")
	fmt.Println("Color Slices1:", myColorsSlice1)	// Color Slices1: [Red Blue Yellow Black White Hello]
//================================================================
	fmt.Println()

	nums := make([]int, 5)
	nums[1] = 11
	fmt.Println(nums)	// [0 11 0 0 0]
	nums = append(nums, 789)
	fmt.Println(nums)	// [0 11 0 0 0 789]

	fmt.Println("LENGTH:", len(nums))
	fmt.Println("LENGTH:", cap(nums))
//================================================================
	fmt.Println()

	myArray2 := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 111}
	mySlice2 := myArray2[:]
	mySlice3 := myArray2[:]

	fmt.Println("O'ZGARISHSIZ")
	fmt.Println("My Array2:", myArray2)	// My Array2: [11 22 33 44 55 66 77 88 99 111]
	fmt.Println("My Slice2:", mySlice2)	// My Slice2: [11 22 33 44 55 66 77 88 99 111]
	fmt.Println("My Slice3:", mySlice3)	// My Slice3: [11 22 33 44 55 66 77 88 99 111]

	mySlice2[0] = 777

	fmt.Println("O'ZGARISH BILAN")
	fmt.Println("My Array2:", myArray2)	//	My Array2: [777 22 33 44 55 66 77 88 99 111]
	fmt.Println("My Slice2:", mySlice2)	//	My Slice2: [777 22 33 44 55 66 77 88 99 111]
	fmt.Println("My Slice3:", mySlice3)	//	My Slice3: [777 22 33 44 55 66 77 88 99 111]

}