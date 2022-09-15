package main

import "fmt"

func main() {

	length, sum := total(10, 20, 30,123,45,59,8,45,26,7,23,56,889)
	fmt.Println("Length:", length, "Total:", sum)

}

func total(nums ... int) (int, int) {
	result := 0
	for _, num := range nums {
		result += num
	}
	return len(nums), result
}