// https://gobyexample.com/closures
package main

import "fmt"

func Increment() func() int{
	i := 0

	return func() int{
		i++
		return i
	}
}

func main() {
	nextInt := Increment()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := Increment()
	fmt.Println(nextInts())
	fmt.Println(nextInts())
	fmt.Println(nextInts())
	
}