package main

import "fmt"

//*======================INTERFACE===========================

// *======================STRUCTS=============================
type Car struct {
	Model, BodyType, Color, TypeOfFuel string
	Motor, Speed, Price, ReleseDate    float64
}

type Phone struct {
	Model, Class, OpSystem, Color string
	ReleseDate, Battery, Cameras  float64
}

type Laptop struct {
	Model, OsSystem, Color, Processor string
	Memory, Storage                   float64
}

type Playstation struct {
	Model, GpuArchitect, GameDisc, CPU string
	Memory, Power, InternalStorage     float64
}

//*======================METHODS============================

func (c Car) Show() {
	fmt.Printf("Model: %s\nBodyType: %s\nColor: %s\nTypeOfFuel: %s\nSpeed: %1.f\nMotor: %1.f\nPrice: %1.f\nReleseDate: %1.f\n",
		c.Model, c.BodyType, c.Color, c.TypeOfFuel, c.Motor, c.Speed, c.Price, c.ReleseDate)
	fmt.Println("===========================")
}

func (p Phone) Show() {
	fmt.Printf("Model: %s\nClass: %s\nOpSystem: %s\nColor: %s\nReleseDate: %1.f\nBattery: %1.f\nCameras: %1.f\n",
		p.Model, p.Class, p.OpSystem, p.Color, p.ReleseDate, p.Battery, p.Cameras)
	fmt.Println("===========================")

}

func (l Laptop) Show() {
	fmt.Printf("Model: %s\nOsSystem: %s\nColor: %s\nProcessor: %s\nMemory: %1.f\nStorage: %1.f\n",
		l.Model, l.OsSystem, l.Color, l.Processor, l.Memory, l.Storage)
	fmt.Println("===========================")
}

func (p Playstation) Show() {
	fmt.Printf("Model: %s\nGpuArchitect: %s\nGameDisc: %s\nCPU: %s\nMemory: %1.f\nPower: %1.f\nInternalStorage: %1.f\n",
		p.Model, p.GpuArchitect, p.GameDisc, p.CPU, p.Memory, p.Power, p.InternalStorage)
	fmt.Println("===========================")
}

func main() {

	car := Car{
		Model:      "Tesla",
		BodyType:   "Sedan",
		Color:      "Black",
		TypeOfFuel: "Electro",
		Motor:      12,
		Speed:      360,
		Price:      80.000,
		ReleseDate: 2022,
	}
	car.Show()

	phone := Phone{
		Model:      "Iphone",
		Class:      "S_Class",
		OpSystem:   "iOS",
		Color:      "Black",
		ReleseDate: 2022,
		Battery:    4323,
		Cameras:    48,
	}
	phone.Show()

	laptop := Laptop{
		Model:     "Lenovo",
		OsSystem:  "RyZen",
		Color:     "Black",
		Processor: "CoreI7",
		Memory:    12,
		Storage:   256,
	}
	laptop.Show()

	play := Playstation{
		Model:           "PS5",
		GpuArchitect:    "Ninth",
		GameDisc:        "GTRX",
		CPU:             "AMD Core:8",
		Memory:          16,
		Power:           12,
		InternalStorage: 8,
	}
	play.Show()

}
