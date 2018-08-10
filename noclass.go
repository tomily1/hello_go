package main

import "fmt"

type Salutations string

type Salutation struct {
	name string
	greeting string
}

const ( // Make declaration brief and succinct
	PI = 3.14
	Language = "Go"
)

const (
	A = iota // iota represents successive untyped integer constants... can be declared once also.
	B = iota
	C = iota
)

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
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)
}
