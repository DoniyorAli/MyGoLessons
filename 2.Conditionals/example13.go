package main   // ASCII TABLE

import "fmt"

func main() {

	// asciiTable1 := 31
	// fmt.Println(string(rune(asciiTable1))) // --> A

	asciiTable2 := 'A'
	fmt.Println(asciiTable2)      // --> 65

	num := 0
	// letter := ""
	fmt.Println("Enter the number from 33 ... to 127!: ")
	fmt.Scanln(&num)
	fmt.Println(string(rune(num)))

	
	// if 97 <= harf && harf >= 122{
	// 	fmt.Println("Small letter!")
	// }else if 65 <= harf && harf <= 90{
	// 	fmt.Println("Big letter!")
	// }
}