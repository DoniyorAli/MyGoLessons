package main
				   // LOGICAL OPERATORS    || and && and !
import "fmt"

func main() {

	// || --> OR (yoki) --> +
	fmt.Println(true || true)   //  1 + 1 = 1 --> true 
	fmt.Println(true || false)  //  1 + 0 = 1 --> true
	fmt.Println(false || true)  //  0 + 1 = 1 --> true
	fmt.Println(false || false) //  0 + 0 = 0 --> false

	// && --> AND (va) --> *
	fmt.Print("\n")
	fmt.Println(true && true)   //  1 * 1 = 1 --> true 
	fmt.Println(true && false)  //  1 * 0 = 0 --> false
	fmt.Println(false && true)  //  0 * 1 = 0 --> false
	fmt.Println(false && false) //  0 * 0 = 0 --> false

	// || and && Examples:
	fmt.Print("\n")
	fmt.Println(true && false || true)   // --> true
	fmt.Println(false || true && false)  // --> false
	fmt.Println(true && true || true)    // --> true
	fmt.Println(false || false && false) // --> false
	fmt.Println(true && false || true)   // --> true
	fmt.Println(false || true && false)  // --> false

	// ! --> 
	fmt.Print("\n")
	fmt.Println(!true)   // --> false
	fmt.Println(!false)  // --> true

}