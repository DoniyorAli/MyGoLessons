package main

import "fmt"

func main() {

	num1 := 15
	num2 := 17
	num3 := 18

// BIG NUMBER
	if num1 < num2{
		num1 = num2
	}
	if num1 < num3{
		num1 = num3
	}
	fmt.Println("Big number: ", num1)

// SMALL NUMBER
	if num1 > num2{
		num1 = num2
	}
	if num1 > num3{
		num1 = num3
	}
	fmt.Println("Small number: ", num1)
//___________________________________________________________

// SECOND EXAMPLE
// BIG NUMBER
// 	if num1 > num2 && num1 > num3{
// 		fmt.Println("Big number:", num1)
// 	}else if num2 > num1 && num2 > num3{
// 		fmt.Println("Big number: ", num2)
// 	}else if num3 > num1 && num3 > num2{
// 		fmt.Println("Big number: ", num3)
// 	}else if num1 == num2 && num1 == num3{
// 		fmt.Println("raqamlar teng!")
// 	}

// //  SMALL NUMBER
// 	if num1 < num2 && num1 < num3{
// 		fmt.Println("Small number:", num1)
// 	}else if num2 < num1 && num2 < num3{
// 		fmt.Println("Small number: ", num2)
// 	}else if num3 < num1 && num3 < num2{
// 		fmt.Println("Small number: ", num3)
// 	}else if num1 == num2 && num1 == num3{
// 		fmt.Println("The numbers equal!")
// 	}

}