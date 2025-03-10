package main

import (
	"fmt"

	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

type IOsystem interface {
	Readlines() ([]string, error)
	WriteResult() error
}

func main() {
	taxRaces := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRace := range taxRaces {
		f := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRace*100))
		priceJob := prices.NewTaxIncludedPriceJob(f, taxRace)
		priceJob.Process(f.InputFilePath)
	}

}
