package main            // FOR 

func main() {

//		declare	   logic    increment
	for num := 0; num < 10; num += 1 {
		println(num) // --> 0,1,2,3,4,5,6,7,8,9
	}
	println("\n")
	for a := 10; a >= 0; a-- {
		println(a) // --> 10,9,8,7,6,5,4,3,2,1,0
	}
	println("\n")
	for i := 10; i < 22; i++ {
		println(i) // 10, 11, 12 13, 14, 15, 16, 17, 18, 19, 20, 21 
	}

}