package main

func CalRev(totalIncome int64, tax int64) (int64, int64) {
	profitBeforeTax := totalIncome

	taxRate := float64(tax) / 100
	taxInNumbers := int64(float64(totalIncome) * taxRate)
	println(taxInNumbers)
	profitAfterTax := totalIncome - taxInNumbers
	println(profitAfterTax)

	return profitBeforeTax, profitAfterTax
}
