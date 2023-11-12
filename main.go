package main

import (
	"cellbesmanos/greetings"
	"fmt"
	"math"
	"math/rand"
)

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
}