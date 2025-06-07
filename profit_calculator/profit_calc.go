package main

import (
	"fmt"
	
)

var revenue float64
var expenses float64
var taxRate float64
var ratio float64

func main() {

	userInput()

	ebt, profit := calculations(revenue, expenses, taxRate)
	
	output(ebt, profit)
	

}

func userInput() {
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
}

func calculations(revenue, expenses, taxRate float64) (ebt float64, profit float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func output(ebt, profit float64) {
	fmt.Println()
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)
}