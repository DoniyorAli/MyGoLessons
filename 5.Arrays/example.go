package main //  ARRAY

import "fmt"

func main() {

	myArray := [5]int{11, 22, 33, 44, 55}
	fmt.Println(myArray)
	fmt.Printf("0.index--->: %v\n", myArray[0])
	fmt.Printf("1.index--->: %v\n", myArray[1])
	fmt.Printf("2.index--->: %v\n", myArray[2])
	fmt.Printf("3.index--->: %v\n", myArray[3])
	fmt.Printf("4.index--->: %v\n", myArray[4])
//========================================================
	fmt.Println()

	var colors [3]string
	fmt.Println(colors)
	colors[0] = "Blue"
	colors[1] = "White"
	colors[2] = "Black"
	fmt.Println(colors)
//========================================================
	fmt.Println()

	namesArray := [5]string{"Ali", "Vali", "G'ani", "Qani", "Dani"}
	fmt.Println(namesArray)
	fmt.Println("3.index name is --->", namesArray[3])
//=========================================================
	fmt.Println()

	numArray := [3]int{}
	fmt.Println(numArray)
	numArray[0] = 12
	numArray[1] = 45
	numArray[2] = 78
	fmt.Println(numArray)
//=========================================================
	fmt.Println()

	var fruits = [5]string{"apple", "banana", "orange", "pear", "grape"}
	fmt.Println(fruits)
//=========================================================
	fmt.Println()

	var numbers = [5]int{1, 3, 5, 7, 9}
	fmt.Println("LENGTH ARRAY:", len(numbers))

}
