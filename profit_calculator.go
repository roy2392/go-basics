//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var revenue float64
//	var expenses float64
//	var taxRate float64
//
//	fmt.Print("Revenue Amount: ")
//	fmt.Scan(&revenue)
//
//	fmt.Print("Expenses Amount: ")
//	fmt.Scan(&expenses)
//
//	fmt.Print("Tax Rate: ")
//	fmt.Scan(&taxRate)
//
//	ebt := revenue - expenses
//	profit := ebt * (1 - taxRate/100)
//	ratio := ebt / profit
//
//	formattedFV := fmt.sprintf("EBT: %.1f\n ", ebt)
//	formattedRFV := fmt.sprintf("Profit: %.1f\n ", profit)
//	// outputs information
//	fmt.Printf(`EBT: %.1f\n
//		Profit: %.1f\n`)
//}
