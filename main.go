package main

import (
	"fmt"
)

// Account struct to hold user details
type Account struct {
	ID           uint
	Name         string
	Balance      float64
	Transactions []string
}

var accounts []Account
var currentID uint

// Function to create a new account
func createAccount(name string, initialBalance float64) Account {
	currentID++
	account := Account{
		ID:           currentID,
		Name:         name,
		Balance:      initialBalance,
		Transactions: []string{fmt.Sprintf("Account created with initial balance of %.2f", initialBalance)},
	}
	accounts = append(accounts, account)
	return account
}

// Function to deposit money into an account
func deposit(accountID uint, amount float64) {
	for i, account := range accounts {
		if account.ID == accountID {
			accounts[i].Balance += amount
			accounts[i].Transactions = append(accounts[i].Transactions, fmt.Sprintf("Deposited %.2f", amount))
			fmt.Printf("Deposited %.2f into account %s. New balance: %.2f\n", amount, account.Name, accounts[i].Balance)
			return
		}
	}
	fmt.Println("Account not found.")
}

// Function to withdraw money from an account
func withdraw(accountID uint, amount float64) {
	for i, account := range accounts {
		if account.ID == accountID {
			if accounts[i].Balance >= amount {
				accounts[i].Balance -= amount
				accounts[i].Transactions = append(accounts[i].Transactions, fmt.Sprintf("Withdrew %.2f", amount))
				fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, account.Name, accounts[i].Balance)
			} else {
				fmt.Println("Insufficient balance.")
			}
			return
		}
	}
	fmt.Println("Account not found.")
}

// Function to display account details
func displayAccount(accountID uint) {
	for _, account := range accounts {
		if account.ID == accountID {
			fmt.Printf("\nAccount ID: %d\nName: %s\nBalance: %.2f\nTransactions:\n", account.ID, account.Name, account.Balance)
			for _, transaction := range account.Transactions {
				fmt.Println(transaction)
			}
			return
		}
	}
	fmt.Println("Account not found.")
}

// Function to display menu and handle user actions
func main() {
	for {
		// Display menu options
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. View Account Details")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var initialBalance float64
			fmt.Print("Enter account holder's name: ")
			fmt.Scan(&name)
			fmt.Print("Enter initial balance: ")
			fmt.Scan(&initialBalance)
			account := createAccount(name, initialBalance)
			fmt.Printf("Account created! Account ID: %d\n", account.ID)
		case 2:
			var accountID uint
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			deposit(accountID, amount)
		case 3:
			var accountID uint
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			withdraw(accountID, amount)
		case 4:
			var accountID uint
			fmt.Print("Enter account ID to view: ")
			fmt.Scan(&accountID)
			displayAccount(accountID)
		case 5:
			fmt.Println("Exiting banking app...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
