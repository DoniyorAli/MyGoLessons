package main

import (
	"fmt"
	"strconv"
)

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

//* INTERFUNCTION
func CarExecute(c Carface) {
	fmt.Println("Car Info: " + c.Information())
	message := ""
	isRun := c.Run()
	if isRun == true {
		message = "Car is working"
	}else {
		message = "Car is not working!"
	}
	fmt.Println(message)

	msg := ""
	isStop := c.Stop()
	if isStop == true {
		msg = "Stopped!"
	}else {
		msg = "Going!"
	}
	fmt.Println(msg)
}

type Car struct {
	Brand, Model, Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}
//* FERRARI STRUCT============================
type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (f Ferrari) Information() string {
	r := f.Brand + "\n" + f.Color + "\n" + f.Model + "\n" + strconv.Itoa(f.Speed) + "\n" + strconv.FormatFloat(f.Price, 'f', -1, 64)+" million $"
	a := "Yes"
	if f.Special == true {
		fmt.Println(a)
	}
	return r
}

//* Mercedes STRUCT============================
type MercedesBens struct {
	Car
	SpecialProduction
}

func (_ MercedesBens) Run() bool {
	return true
}

func (_ MercedesBens) Stop() bool {
	return true
}

func (f MercedesBens) Information() string {
	r := f.Brand + "\n" + f.Color + "\n" + f.Model + "\n" + strconv.Itoa(f.Speed) + "\n" + strconv.FormatFloat(f.Price, 'f', -1, 64)+" million $"
	a := "Yes"
	if f.Special == true {
		fmt.Println(a)
	}
	return r
}

func main() {

	f1 := new(Ferrari)
	f1.Brand = "F45"
	f1.Model = "Ferrari F1"
	f1.Color = "Red"
	f1.Speed = 360
	f1.Price = 1.2
	f1.Special = true
	// fmt.Println(f1.Information())

	fmt.Println("==================")

	m1 := new(MercedesBens)
	m1.Brand = "Maybach"
	m1.Model = "S 222 class"
	m1.Color = "Black"
	m1.Speed = 300
	m1.Price = 1
	m1.Special = false
	// fmt.Println(m1.Information())

	CarExecute(f1)
	CarExecute(m1)
}