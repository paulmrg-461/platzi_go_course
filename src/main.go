package main

import (
	"fmt"
)

func main() {

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi: ", pi)
	fmt.Println("Pi 2: ", pi2)

	base := 12          // Declare variable, assign value and infer data type
	var altura int = 14 // Declare variable, assign value and assign data type
	var area int        // Declare variable, and assign data type

	area = (base * altura) / 2

	fmt.Println("Area triangulo: ", base, "m X", altura, "m =", area)

	//Zero values
	//Go set default values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
