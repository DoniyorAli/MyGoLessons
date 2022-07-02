package main

import "fmt"

func main() {

	asciiTable := 125
	// fmt.Println(string(rune(asciiTable))) // --> A
	// fmt.Println(rune(asciiTable))         // --> 65
	box := string(rune(asciiTable))
	fmt.Println(box)
	

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