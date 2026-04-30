package main

import (
	"fmt"

	"example.com/greetings"
)
func main() {
	fmt.Println("Hello World")
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
