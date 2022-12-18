package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {

	fmt.Println(Vertex{1, 2})
	v := Vertex{11, 22}
	adress := &v
	adress.X = 1e9
	fmt.Println(v)
}