package main

// This is the last excercire
// Goal
// 1) Validate user input
// 	=> show error message and exit if invalid input is provide
// 	- No negative numbers
// 	- Not 0
// 2) Store calculated result into file

import (
	"errors"
	"fmt"
	"os"
)

func getUserInput(inforText string) (float64, error) {
	var userInput float64
	fmt.Print(inforText)
	fmt.Scanln(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid userInput!")
	}
	return userInput, nil
}

func WriteIntoFile(revernue, expenses, taxRate float64) {
	data := fmt.Sprintf("Revernue: %.0f\nExpenses: %.0f\nTax: %.0f", revernue, expenses, taxRate)
	os.WriteFile("userData.txt", []byte(data), 0644)
}

func main() {
	revernue, err := getUserInput("Input revernue: ")
	if err != nil {
		fmt.Print("ERROR")
		panic(err)
	}
	expenses, err := getUserInput("Input expenses: ")
	if err != nil {
		fmt.Print("ERROR")
		panic(err)
	}
	taxRate, err := getUserInput("Input taxRate: ")
	if err != nil {
		fmt.Print("ERROR")
		panic(err)
	}
	fmt.Println(revernue)
	fmt.Println(expenses)
	fmt.Println(taxRate)

	WriteIntoFile(revernue, expenses, taxRate)
}

// revenue
// expenses
// taxRate

// ebt, profit, ratio
