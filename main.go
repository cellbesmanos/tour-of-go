package main

import (
	"cellbesmanos/greetings"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// this is a custom package
	greetings.Hello("test")
}