package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.Open("History.txt")
	checkError(error)
	fmt.Println("FILE:", file.Name())

}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}