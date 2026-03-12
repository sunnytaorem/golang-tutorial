package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.0 // Annual inflation rate in %
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	var futureValue float64

	// You can also declare variables in a single line:
	// var investmentAmount, expectedReturnRate, years, futureValue float64

	// You can also inherit variable types from the assigned value:
	// investmentAmount := 0.0
	// expectedReturnRate := 0.0
	// years := 0.0
	// futureValue := 0.0

	// Get user input
	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scanln(&investmentAmount) // The & symbol (pointer) is used to pass the address of the variable to the Scanln function, allowing it to modify the variable's value.
	// fmt.Print("Enter the expected annual return rate (in %): ")
	outputText("Enter the expected annual return rate (in %): ")
	fmt.Scanln(&expectedReturnRate)
	// fmt.Print("Enter the number of years: ")
	outputText("Enter the number of years: ")
	fmt.Scanln(&years)

	// Calculate future value
	// futureValue = calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	// Display the result
	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
	fmt.Printf("Future value of the investment adjusted for inflation: %.2f\n", futureRealValue)

	// Store the future value in a variable for later use
	formattedFV := fmt.Sprintf("Stored future value for later use: %.2f\n", futureValue)
	formattedRV := fmt.Sprintf("Stored future real value for later use: %.2f\n", futureRealValue)
	fmt.Print(formattedFV)
	fmt.Print(formattedRV)
}

func outputText(text string) {
	fmt.Println(text)
}

// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
// 	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// }

// func calculateFutureRealValue(futureValue, inflationRate, years float64) float64 {
// 	return futureValue / math.Pow(1+inflationRate/100, years)
// }

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}

// Alternative way to return multiple values
func calculateFutureValuesAlt(investmentAmount, expectedReturnRate, inflationRate, years float64) (fv float64, frv float64) {
	// In this case, you dont have to declare the variables fv and frv inside the function, as they are already declared in the function signature. You can directly assign values to them and return without specifying them in the return statement.
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv // You can also return the values without specifying the variable names, as they are already declared in the function signature.
	// return
}

// You can also use a struct to return multiple values
type FutureValues struct {
	FutureValue     float64
	FutureRealValue float64
}

func calculateFutureValuesStruct(investmentAmount, expectedReturnRate, inflationRate, years float64) FutureValues {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return FutureValues{FutureValue: fv, FutureRealValue: frv}
}
