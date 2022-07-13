package main

import "fmt"

func main() {

	a := 10
	b := 20

	total := a + b
	fmt.Println(total)

	modul := a % b    // Harqanday kichik sonni katta songa modulini olsak to'g'ridan 
	fmt.Println(modul)  // --> 10 // to'g'ri kichik sonning o'zi o'tadi

	num1 := 63
	num2 := 10

	module := num1 % num2
	fmt.Println(module)
//______________________________________________________________________

	num := 0
	fmt.Println(num)
	num += 10  // num = num + 10 --> 10
	fmt.Println(num)
	num -= 10  // num = num - 10 --> 0 
	fmt.Println(num)
	num *= 10  // num = num * 10 --> 0
	fmt.Println(num)
	num /= 10  // num = num / 10 --> 0
	fmt.Println(num)
//_______________________________________________________________________

	box := 10
	box -= 5  // box = 10 - 5 --> 5
	box += 4  // box = 5 + 4  --> 9
	box *= 2  // box = 9 * 2  --> 18
	box /= 3  // box = 18 / 3 --> 6
	fmt.Println(box)
//_______________________________________________________________________

	number := 10
	number++            // number += 1  bu ham bo'ladi
	fmt.Println(number) // --> 11


	number--            // number -= 1  bu ham bo'ladi
	fmt.Println(number) // --> 10
//______________________________________________________________________

	number += 5        
	fmt.Println(number)  // --> 15

	number -= 3		  
	fmt.Println(number)  // --> 12

	number /= 2
	fmt.Println(number)  // --> 6

	number *= 3
	fmt.Println(number)  // --> 18

}