package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRaces := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)
	result[1.1] = []float64{2.2, 2.3}
	for _, taxRace := range taxRaces {
		taxIncludedPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = price * (1 + taxRace)
		}
		result[taxRace] = taxIncludedPrice
	}

	for _, taxRace := range taxRaces {
		for _, price := range prices {
			fmt.Println("result[taxRace]", result[taxRace], price, " ", taxRace)
		}
		fmt.Println()
	}

}
