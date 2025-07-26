package main

import (
	"fmt"

	"go-journey-from-scratch/01-bank-simulator/bank"
)

func main() {
	myAccount := bank.NewAccount(2000.00)
	fmt.Printf("Initial balance: $%.2f\n", myAccount.Balance())

	fmt.Println("Adding 1.000 money to the account")
	myAccount.Deposit(1000.00)
	fmt.Printf("Current balance after adding 1.000 money to the account: $%.2f\n", myAccount.Balance())

	fmt.Println("Withdrawing 500 money from the account")
	myAccount.Withdraw(500.00)
	fmt.Printf("Current balance after withdrawing 500 money from the account: $%.2f\n", myAccount.Balance())
}
