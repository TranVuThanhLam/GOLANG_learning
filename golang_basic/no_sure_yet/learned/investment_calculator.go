package main

// What we learn in this:
// - Var, :=
// fmt.Print, fmt.Scan, fmt.Printf("%.1f",0.111)
import (
	"fmt"
	"math"
)

func investment() {
	var investmentAmount = 1000
	var expectReturnRate = 5.5
	var years = 10
	fmt.Scan(investmentAmount)
	var futureValue = float64(investmentAmount) * (math.Pow(1+expectReturnRate/100, float64(years)))
	fmt.Println(futureValue)
}
