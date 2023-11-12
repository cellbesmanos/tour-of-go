package main

import (
	"cellbesmanos/greetings"
	"fmt"
	"math"
	"math/rand"
)

// types are automatically assigned here based on the initializer
var c, python, java = true, false, "test"

// this is a naked return, notice the return type has variables
// this will automatically finds the variables within the function
// and return their values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	
	// notice that removing this part will return y as 0 and not nil
	y = sum - x
	return
}

// reminds me of Python
func swapValues(x, y string) (string, string) {
	return y, x
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// this is a custom package
	greetings.Hello("test")

	fmt.Println("\n", add(1, 2))

	a := "World"
	b := "Hello"
	a, b = swapValues(a, b)

	fmt.Println(a, b, "!")	

	fmt.Println(split(17))

	var sample = 100
	fmt.Println(sample, c, python, java)

	// you cannot use the shorthand outside a func, since it does not start with a keyword
	// unlike the var declaration 
	shortHand := "This is a shorthand variables assignment and declaration!"
	fmt.Println(shortHand)

	// declared variables without explicit values will get their "zero" value
	// 0, false, and empty string for numeric, bool, and string types respectively
	var num int
	var booool bool 
	var str string
	fmt.Println(num, booool, str)

	// assignment between different types requires explicit conversion
	toBeFloat := -64
	aFloat := float64(toBeFloat)
	fmt.Println(aFloat)

	// constants cannot use the := syntax
	const PI = 3.14
	fmt.Println(PI)
}