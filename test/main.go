package main

import "fmt"

func main() {
	var x interface{} = "Golang"
	if v, ok := x.(int); ok {
		fmt.Println(v)
	} else {
		fmt.Println("Not an int")
	}
}
