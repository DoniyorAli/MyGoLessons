package main

import "fmt"

func main() {
// ! Variadic functionga foydalanuvchidan qiymat qabul qilip boladimi 
	// var box int
	// for i := 0; i <= 10; i++ {
	// 	fmt.Scanf("%d", &box)
	// 	append()
	// }

	fmt.Println("Total compute:", computeNums())

}

func computeNums(nums ... int) (total int) {
	for _, value := range nums {
		total += value
	}
	return
}
//? QUESTION