package main

import (
	"fmt"
	"log"
	"strconv"
)

func printMessage(message string, iterarions int) {
	for i := 0; i < iterarions; i++ {
		fmt.Printf("Message: %s, Iteration: %d of %d\n", message, i+1, iterarions)
	}
}

func forWhile(condition int) {
	counter := 0
	for counter <= condition {
		fmt.Printf("Current: %d of: %d\n", counter, condition)
		counter++
	}
}

func sum2Integers(a, b int) int {
	return a + b
}

func applyCommonOperators(a, b int) (addition, subtraction, multiplication, division int) {
	return a + b, a - b, a * b, a / b
}

func comparePairNumber(number int) bool {
	return number%2 == 0
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return (fibonacci(n-1) + fibonacci(n-2))
	}
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

	forWhile(10)

	// Convert text to number
	value, err := strconv.Atoi("17")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Converted value = %d\n", value)

	// Print if number is pair
	for i := 0; i <= 10; i++ {
		if comparePairNumber(i) {
			fmt.Printf("Number %d: is pair\n", i)
			continue
		}
		fmt.Printf("Number %d: is odd\n", i)
	}

	// Print if number is pair using switch sentence
	for i := 0; i <= 10; i++ {
		switch condition := i % 2; condition {
		case 0:
			fmt.Printf("Number %d: is pair\n", i)
		default:
			fmt.Printf("Number %d: is odd\n", i)
		}
	}

	// Print if number is pair using switch without condition sentence
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci: %d\n", fibonacci(i))
	}

	for i := 0; i <= 10; i++ {
		switch {
		case i >= 0 && i <= 5:
			fmt.Printf("Number %d: is in 0 - 5 range\n", i)
		case i > 5 && i <= 10:
			fmt.Printf("Number %d: is in 5 - 10 range\n", i)
		default:
			fmt.Printf("Number %d: \n", i)
		}
	}

	// defer keyword execute this function on final  current 1. Hello 2. ----, 3. World
	// defer to close connections, streams... (Similar to dispose)
	defer fmt.Println("World")
	fmt.Println("Hello")
	fmt.Println("------")

	// Arrays (Immutable)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slices (To left of : is inclusive and to right of : is exclusive)
	slice := []int{2, 0, 2, 4, 1, 17}
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Printf("Slice[0] : %d\n", slice[0])
	fmt.Printf("Slice[0] - [3]	: %d\n", slice[:3])
	fmt.Printf("Slice[2] - [4]	: %d\n", slice[2:4])
	fmt.Printf("Slice[4] - []	: %d\n", slice[4:])

	// Append
	slice = append(slice, 18)
	fmt.Println(slice, len(slice), cap(slice))
	slice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slice = append(slice, slice2...)
	fmt.Println(slice, len(slice), cap(slice))
}
