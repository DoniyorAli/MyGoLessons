package main

import "fmt"

func main() {

	aboutI("Alixan", "Farhodov", 26)
	fmt.Println(aboutYou("Maryam", "Farhodova", 18))
	fmt.Println(aboutYou("Mustafo", "Komiljonov", 3))

}

//!=========================Functions========================================

func aboutI(name, fname string, age int) {
	fmt.Println("My name is", name, "first name is", fname, "my age is", age)
}
//*==========================================================================

func aboutYou(name, fname string, age int) (string, string, int) {
	return name, fname, age
}
//*==========================================================================

func aboutChild(name string, fname string, age int) (string, string, int) {
	return name, fname, age
}
