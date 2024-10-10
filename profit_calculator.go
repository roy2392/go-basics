package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expanses := getUserInput("Expanses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expanses, taxRate)

	fmt.Print(ebt)
	fmt.Print(profit)
	fmt.Print(ratio)
}

func getUserInput(infoText string) float64 {
	var userValue float64
	fmt.Print(infoText)
	fmt.Scan(&userValue)
	return userValue
}

func calculateFinancials(revenue, expanses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expanses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}
