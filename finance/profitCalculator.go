package main

import (
	"errors"
	"fmt"
	"os"
)

const financeResults = "financial_results.txt"

func main() {
	// Get user input
	revenue, err := getUserInput("Enter the total revenue: ")
	if err != nil {
		fmt.Println(err)
		return //this will exit the main function immediately.
		// panic(err) // panic will print the error message and stack trace before exiting the program, which can be helpful for debugging.
	}
	expenses, err := getUserInput("Enter the total expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter the tax rate (in %): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateEarnings(revenue, expenses, taxRate)
	saveResultsToFile(earningsBeforeTax, profit, ratio)

	// Display the result
	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Profit ratio (before tax / after tax): %.2f\n", ratio)
}

func getUserInput(inputText string) (float64, error) {
	var userInput float64
	fmt.Println(inputText)
	fmt.Scanln(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid input. Please enter a positive value.")
	}
	return userInput, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func saveResultsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio (EBT / Profit): %.2f\n", ebt, profit, ratio)
	os.WriteFile(financeResults, []byte(results), 0644)
}
