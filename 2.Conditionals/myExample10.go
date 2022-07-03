package main

import "fmt"

func main() {

	// asciiTable1 := 31
	// fmt.Println(string(rune(asciiTable1))) // --> A

	// asciiTable2 := 'A'
	// fmt.Println(asciiTable2)         // --> 65

	num := 0
	// letter := ""
	fmt.Println("Enter the number from 33 ... to 127!: ")
	fmt.Scanln(&num)
	fmt.Println(string(rune(num)))
	
	// fmt.Println("Enter the Alphabet letters: ")
	// fmt.Scanln(&letter)
	// fmt.Println(letter)

	// asciiT := 0
	// fmt.Println("Son kiriting: ")
	// fmt.Scanln(&asciiT)
	// character := rune(asciiT)
	// fmt.Printf("ASCII demical %d --> %c\n", asciiT, character)
	

	// if 97 <= harf && harf >= 122{
	// 	fmt.Println("Small letter!")
	// }else if 65 <= harf && harf <= 90{
	// 	fmt.Println("Big letter!")
	// }
}