package main    // MULTIPLE TABLE

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")

	for i := 1; i <= 10; i++ {
		for j := 6; j <= 10; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Print("\n")
	}
}

