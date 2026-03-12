package main

import (
	"fmt"
	"go-project/bank/fileops" // You need to include the module path to import the fileops package.

	"github.com/Pallinder/go-randomdata"
)

// You can search for packages available in Go using this link: https://pkg.go.dev/

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile) // only works if there is a balance.txt file with a valid float value in it. If the file doesn't exist or contains invalid data, the balance will be initialized to 0.00.
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't continue, sorry!")
		return // You can use return to exit the main function immediately, but using panic will print the error message and stack trace before exiting the program, which can be helpful for debugging.
	}
	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber()) // This is just a fun way to generate random data using the go-randomdata package. You can remove this line if you don't need it.
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice (1-8): ")
		fmt.Scanln(&choice)

		fmt.Printf("You chose: %d\n", choice)

		// Handle the user's choice using a switch statement. You can also use if-else statements, but switch is often more concise and easier to read when dealing with multiple discrete cases.
		switch choice {
		case 1:
			fmt.Println("Opening an account...")
		case 2:
			fmt.Println("Closing an account...")
		case 3:
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 4:
			fmt.Println("Viewing account details...")
		case 5:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scanln(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please enter a positive value.")
				// return // You can use return to exit the main function immediately, but using continue will skip the rest of the loop and prompt the user again without exiting the program.
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Printf("Balance updated! Your new balance after deposit is: $%.2f\n", accountBalance)
		case 6:
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scanln(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Please enter a positive value.")
				// return // You can use return to exit the main function immediately, but using continue will skip the rest of the loop and prompt the user again without exiting the program.
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds! Withdrawal amount exceeds current balance.")
				continue
			}
			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Printf("Balance updated! Your new balance after withdrawal is: $%.2f\n", accountBalance)
		case 7:
			fmt.Println("Transferring money...")
		case 8:
			fmt.Println("Exiting the bank...")
			fmt.Println("Thank you for using Go Bank. Have a nice day!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
			// return
		}
	}
}
