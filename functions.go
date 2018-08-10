package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func main() {
	var s = Salutation {}
	s.name = "Bob"
	s.greeting = "Hello"

	Greet(s)
	CreateGreet(s)

	_, alternate := CreateMessageDouble(s.name, s.greeting)

	// fmt.Println(message)
	fmt.Println(alternate)
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func CreateGreet(salutation Salutation) {
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func CreateMessageDouble(name, greeting string) (string, string) {
	return greeting + " " + name, "HEY!" + " " + name
}