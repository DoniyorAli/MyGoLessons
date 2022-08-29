package main          // CONST  ----> O'ZGARMAS

const name = "Alixan"

func main() {

	var a = "A"
	a = "B"
	println(a)

	const box string = "Hello"
	// box = "Salom"   ---> XATO

	println(name)
	println(box)

	// var num1 = 11
	// const num2 = num1    // BU XATO
	// println(num2)

	const box1 = 22
	const box2 = box1
	println(box2)
}