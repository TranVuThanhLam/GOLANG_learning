package main

// What we learn in this:
// Write, read file, convert to string, byte to string
// Catching error by err and panic()

import (
	"fmt"

	"example.com/my-app/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func bank() {
	balance, err := fileops.GetFloatFromFile(fileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
	}
	for {
		bank_menu()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("balance: ", balance)
		} else if choice == 2 {
			var Deposit float64
			fmt.Print("Deposit you want: ")
			fmt.Scanln(&Deposit)

			balance += Deposit
			fmt.Println("Your new balance: ", balance)
			fileops.WriteFloatToFile(balance, fileName)
		} else if choice == 3 {
			var Withdraw float64
			fmt.Print("Withdraw you want: ")
			fmt.Scanln(&Withdraw)
			balance -= Withdraw
			fmt.Println("Your new balance: ", balance)
			fileops.WriteFloatToFile(balance, fileName)
		} else if choice == 4 {
			fmt.Println("Exit")
			break
		}
	}

	fmt.Println("Thank for using")
	fmt.Println(randomdata.PhoneNumber())
}
