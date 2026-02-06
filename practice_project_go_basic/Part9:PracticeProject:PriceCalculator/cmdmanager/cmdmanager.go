package cmdmanager

import "fmt"

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (cm CMDManager) ReadLines() ([]string, error) {

	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
