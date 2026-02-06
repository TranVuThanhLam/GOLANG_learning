package main

import "fmt"

func main() {
	// prices := []float64{10.99, 8.99}
	// fmt.Println(prices[0:1])
	var prices []float64
	prices = append(prices, 12.49, 15.99)
	fmt.Println(prices)

	prices = append(prices, 12.49)
	fmt.Println(prices)

	prices = append(prices, 12.49, 15.99)
	fmt.Println(prices)
	fmt.Println(cap(prices))
}

// type User struct {
// 	Name  string
// 	Email string
// 	Age   int
// }

// user := User{
// 	Name:  "Alice",
// 	Email: "alice@example.com",
// 	Age:   30,
// }
// fmt.Printf("User: %+v\n", user)

// balance := 3000.2
// 	balanceText := fmt.Sprint(balance)
// 	fmt.Printf("balance: %f\nbalanceText: %s\nbyte balance text: %v", balance, balanceText, []byte(balanceText))
