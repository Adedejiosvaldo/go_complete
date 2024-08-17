package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 0.1
	var investedAmount float64
	var years, returnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investedAmount)

	fmt.Print("Amount of Years it will be invested for : ")
	fmt.Scan(&years)

	fmt.Print("What is the return rate: ")
	fmt.Scan(&returnRate)

	futureValue := float64(investedAmount) * math.Pow(1+returnRate/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value is ", futureValue)
	fmt.Println("Future Real value is ", futureRealValue)
}
