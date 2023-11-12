package greetings

import "fmt"

// this is an exported name because it starts with a capital letter
func Hello(name string) {
	hello(name)
}

// this is an unexported name because it starts with a small letter
func hello(name string) {
	fmt.Printf("Hello, %s!", name);
}