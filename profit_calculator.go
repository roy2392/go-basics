package main

import "fmt"

func main() {
	revenue := inputRevenue()
	expanses := inputExpanses()
	taxRate := inputTaxRate()

	ebt := calculateEBT(revenue, expanses)
	profit := calculateProfit(wbt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Print(ebt)
	fmt.Print(profit)
	fmt.Print(ratio)
}

func inputRevenue() float64 {
	var revenue float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	return revenue
}

func inputExpanses() float64 {
	var expanses float64
	fmt.Print("Expanses: ")
	fmt.Scan(&expanses)
	return expanses
}

func inputTaxRate() float64 {
	var taxRate float64
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	return taxRate
}

func calculateEBT(revenue, expanses float64) float64 {
	return revenue - expanses
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}
