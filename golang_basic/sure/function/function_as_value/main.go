package main

import "fmt"

// function as a value
// and it can be Execute and also can be use as parameter

// we use custome type to keep it short

type funcFloat func(float64) float64

func main() {
	number := 7.5
	doubled := transformNumber(number, double)
	tripled := transformNumber(number, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumber(value float64, transform funcFloat) float64 {
	return transform(value)
}

func double(value float64) float64 {
	return value * 2
}
func triple(value float64) float64 {
	return value * 3
}
