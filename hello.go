package main
import "fmt"

func main() {
	/*
	 * Types declaration
	 */

	 // Basic
		var message string
		message = "Hello, Go"

		var a, b, c int = 1, 2, 3

	// Alternative
		var messages = "Hello Go"
		var d, e, f = 1, false, "string"

	// 2nd Alternative
		messagess := "Hello Go"
		g, h, i := 1, 2, true

	// Pointers

	var greeting *string = &message

	fmt.Println(message, a, b, c)
	fmt.Println(messages, d, e, f)
	fmt.Println(messagess, g, h, i)
	fmt.Println(greeting, message) // prints memory address and actual string
	fmt.Println(*greeting, message) // prints pointed message and actual string
}
