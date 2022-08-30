package main      // BREAK

import "fmt"

func main() {

	for i := 0; i < 10; i += 1 {
		if i == 7 {
			break
		}
		fmt.Println(i)  // --> 0,1,2,3,4,5,6
	}
//================================
	sum := 0
	for j := 1; j < 5; j++ {
		if j == 4 {
			break  // ---> 4==4 this break
		}
		sum += j  // 0+1...1+2...3+3 = 6 
	}
	fmt.Println("total:", sum) // 6
}