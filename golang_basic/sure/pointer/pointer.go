package main

import "fmt"

func main() {
	var a = 3
	// var aPointer *int = &a
	fmt.Println(a)
	change(&a)
	fmt.Println(a)
}

func change(a *int) {
	*a = 5
}
