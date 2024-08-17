package main

import (
	"fmt"
	"math"
)

func main() {
	investedAmount, years, returnRate := 1000.0, 10.0, 5.5

	futureValue := float64(investedAmount) * math.Pow(1+returnRate/100, float64(years))
	fmt.Println("Future value is ", futureValue)
}
