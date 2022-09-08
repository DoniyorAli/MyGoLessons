package main

import "fmt"

func main() {

	states := make(map[string]string)
	fmt.Println("STATES:", states)	// STATES: map[]

	states["TAS"] = "Tashkent"
	states["KAS"] = "Kashkadarya"
	states["BUH"] = "Bukhara"
	fmt.Println("STATES:", states)				// STATES: map[BUH:Bukhara KAS:Kashkadarya TAS:Tashkent]
	fmt.Println("STATE ONE:", states["TAS"])	// STATE ONE: Tashkent
	fmt.Println("STATE ONE:", states["BUH"])	// STATE ONE: Bukhara

	buh := states["BUH"]
	fmt.Println("KAS ME:", buh)	// KAS ME: Bukhara

	delete(states, "BUH")
	fmt.Println("STATES:", states)	// STATES: map[KAS:Kashkadarya TAS:Tashkent]

	fmt.Println()

	for k, v := range states {
		fmt.Printf("KEY:%v. VALUE:%v\n",k, v)
	}
	// KEY:TAS. VALUE:Tashkent
	// KEY:KAS. VALUE:Kashkadarya

}