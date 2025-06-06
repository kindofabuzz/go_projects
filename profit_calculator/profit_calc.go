package main

import (
	"fmt"
	
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64
	// var earningsBeforeTax float64
	// var earningsAfterTax float64 
	var ratio float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt  := (revenue - expenses)
	profit := (ebt * (1 - taxRate/100))
	ratio = (ebt / profit)
	
	fmt.Println()
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)

}