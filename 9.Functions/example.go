package main

import "fmt"

func main() {

	helloGo()
	myName("Alixan")
	fmt.Println(myFname("Farkhodov"))
	addNumber(1,2)
	fmt.Println(add(2,3))

}

//!========================Functions==========================

func helloGo() {
	fmt.Println("Hello GO")
}
//*===========================================================

func myName(name string) {
	fmt.Println("Hi", name)
}
//*===========================================================

func myFname(name string) string {
	return name
}

//*===========================================================

func addNumber(n1, n2 int) {
	fmt.Println(n1 + n2)
}
//*===========================================================

func add(n1, n2 int) int{
	return n1 + n2
}