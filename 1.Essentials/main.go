package main

import (
	"fmt"
)

func main() {
	const inflationRate = 0.1
	// var investedAmount float64
	// var years, returnRate float64

	var totalIncome int64
	var tax int64

	// fmt.Print("Investment Amount: ")
	// fmt.Scan(&investedAmount)

	fmt.Print("Total Income: ")
	fmt.Scan(&totalIncome)
	// fmt.Print("Amount of Years it will be invested for : ")
	// fmt.Scan(&years)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&tax)

	// futureValue := float64(investedAmount) * math.Pow(1+returnRate/100, float64(years))

	profitBefore, profitAfter := CalRev(totalIncome, tax)

	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Profit Before Tax", profitBefore)
	fmt.Println("Profit After Tax", profitAfter)
	// Allows us to choose what decimal point we want
	fmt.Printf("The Tex is %.1f\n", float64(profitAfter)) // fmt.Println("Future Real value is ", futureRealValue)
}
