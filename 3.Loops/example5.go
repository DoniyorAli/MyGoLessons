package main    // RANGE

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
//==================================================
	fmt.Print("=======\n")
	
	fructs := map[string]string{"a":"apple", "b":"banana", "g":"grape", "k":"kiwi", "p":"pear", "ch":"cherry"}
	for key, value := range fructs {
		fmt.Println(key,"--->", value)
	}

}