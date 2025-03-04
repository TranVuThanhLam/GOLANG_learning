package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}
