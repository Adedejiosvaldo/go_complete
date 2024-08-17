package main

import (
	"fmt"
	"math"
)

func main() {

	var investedAmount = 1000
	var returnRate = 0.1
	var years = 7

	var futureValue = float64(investedAmount) * math.Pow(1+returnRate/100, float64(years))
	fmt.Println("Future value is ", futureValue)
}
