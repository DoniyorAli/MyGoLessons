package main

import (
	"fmt"
	"os"
)

func main() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	/*uName := os.Getenv("USER")
	fmt.Println("Username: " + uName)

	homePath := os.Getenv("HOME")
	fmt.Println("HomePath: " + homePath)
	*/

	goPath := os.Getenv("_")
	fmt.Println("Go Path: " + goPath)
}
