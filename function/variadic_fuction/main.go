package main

// variadic function
// and splitting slice into parameter value

import "fmt"

func main() {
	var number []any = []any{1, "ahihi", 5}
	display(number)
}

func display(number ...any) {
	for _, val := range number {
		fmt.Print(val)
	}
}
