package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())

	x := "him"

	x = "her"

	fmt.Println("I am " + x)

	fmt.Println("Getting value from myFunc: " + myFunc())
}

func myFunc() string {
	return "hello"
}
