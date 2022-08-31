package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.Open("TheBook.txt")
	if error != nil {
		fmt.Println("There is not TheBook (error)!")
	}else {
		fmt.Println("There is a file:", file.Name())
	}
//=============================================================
	fmt.Println("")

	file, err := os.Open("ProBook.txt")    // There is not ProBook file in the folder
	if err != nil {
		fmt.Println("There is not ProBook (error)!")
	}else {
		fmt.Println("There is a file:", file.Name())
	}
}