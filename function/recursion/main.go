package main

import "fmt"

func main() {
	number := 5
	fmt.Print(factorial(number))
}

func factorial(value int) int {
	if value == 1 {
		return value
	}
	return value * factorial(value-1)
}
