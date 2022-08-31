package main

import "fmt"

func main() {

	myArray := [...]int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println("Elements--->", len(myArray))                    // 9
	fmt.Println("Last element number:", myArray[len(myArray)-1]) // 99
	// BUT! ---> myArray[9] = 14  it is wrong (error)
//=======================================================================
	fmt.Println()

	myCars := [...]string{"MercedesBens", "BMW", "Jaguar", "Hyundai", "Bentley"}
	fmt.Println(myCars)
	fmt.Println("Length:", len(myCars))
	myCars[2] = "Porsche"
	fmt.Println(myCars)
//=======================================================================
	fmt.Println()

	for i := 0; i < len(myCars); i += 1 {
		fmt.Println(i, "--->", myCars[i])
	}
//=======================================================================
	fmt.Println()

	total := 0
	numbers := [...]int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], "+")
		total += numbers[i]
	}
	fmt.Println("---> Total:",total)

	
}
