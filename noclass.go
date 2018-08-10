package main

import "fmt"

type Salutations string

type Salutation struct {
	name string
	greeting string
}

func main() {
	/*
	 * Types declaration
	 */

	 var message Salutations = "Hello WOrld!"

	 var s = Salutation{ "Joe", "Hello" } // equivalent of s.name = Joe, s.greeting = Hello

	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.greeting)

	fmt.Println(message)
}
