package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 
	var expectedReturnRate float64
	var years float64

	// fmt.Print("Investment amount:  ")
	outputText("Investment amount:  ")
	fmt.Scan(&investmentAmount)
	
	// fmt.Print("For how many years?:  ")
	outputText("For how many years?:  ")
	fmt.Scan(&years)

	// fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ") //call outputText function
	fmt.Scan(&expectedReturnRate)

	
	
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV :=	fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

	fmt.Printf(`this is just to
	show multiline output`)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years) 
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv 
	// return
}