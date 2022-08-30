package main // CONTINUE

import "fmt"

func main() {

	for num := 1; num < 10; num++ {
		if num == 3 {
			continue // 3 == 3 continue (tepadan qaytadan davom et !)
		} else if num == 8 {
			break // 8 == 8 break
		}
		fmt.Println(num) // 1, 2, 4, 5, 6, 7
	}
	//==========================================================================

	var users = []string{"Ali", "Dani", "John", "Leyla"}
	for i, v := range users {
		if i == 1{
			continue
		}
		fmt.Println("Value===>", v)
	}
}
