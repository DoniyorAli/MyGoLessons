package main

import "fmt"

func main() {

	name1, name2 := swap("Ali", "Khan")
	fmt.Println("New names:", name1, name2)

}

func swap(a, b string) (string, string){
	fmt.Println("Old names:", a, b)
	return b, a
}