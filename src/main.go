package main

import (
	"fmt"
)

func printMessage(message string, iterarions int) {
	for i := 0; i < iterarions; i++ {
		fmt.Printf("Message: %s, Iteration: %d of %d\n", message, i+1, iterarions)
	}
}

func sum2Integers(a, b int) int {
	return a + b
}

func applyCommonOperators(a, b int) (addition, subtraction, multiplication, division int) {
	return a + b, a - b, a * b, a / b
}

func main() {

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi: ", pi)
	fmt.Println("Pi 2: ", pi2)

	base := 12          // Declare variable, assign value and infer data type first assignation
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

	x := 17
	y := 24
	fmt.Println("Suma:", sum2Integers(x, y))

	// Result variable reassignment ':= value', must be '='
	result := x - y
	fmt.Println("Resta:", result)

	result = x * y
	fmt.Println("Multiplicacion:", result)

	result = x / y
	fmt.Println("Division:", result)

	result = x % y
	fmt.Println("Modulo:", result)

	platform_name := "Platzi"
	courses := 500

	// %s: text, %d: numeric, %v any datatype, %T: Get datatype
	fmt.Printf("%s tiene más de %d cursos\n", platform_name, courses)

	fmt.Printf("La variable 'courses' es de tipo : %T\n", courses)

	printMessage("Hello World!", 3)

	// Multiple assignment with multiple return functions
	addition, substraction, multiplication, division := applyCommonOperators(x, y)
	fmt.Printf("Suma: %d, Resta: %d, Multiplicacion: %d, Division: %d\n", addition, substraction, multiplication, division)

	// Multiple assignment with multiple return functions: Select variables to assign from function
	addition1, _, _, division1 := applyCommonOperators(x, y)
	fmt.Printf("Suma: %d, Division: %d\n", addition1, division1)
}
