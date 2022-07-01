package main // OUTPUT  ---  INPUT

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	println("Hello Golang")   //  OUTPUT

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Raqam kiriting: ")
	num1, _ := reader.ReadString('\n')
	fmt.Println(num1)


	fmt.Println("Enter the number: ")
	str, _ := reader.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)
	fmt.Println(value + 100)
}