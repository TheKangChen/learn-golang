package main

import "fmt"

type Account struct {
	ID      int
	Balance float64
}

func transferFunds(amount float64, source *Account, destination *Account) {
	if source.Balance < amount {
		goto Fail
	} else {
		goto Success
	}
Fail:
	fmt.Printf("Insufficient balance to transfer from account %d\n", source.ID)
	return
Success:
	source.Balance -= amount
	destination.Balance += amount
	fmt.Printf("Successfully transered %g from account %d to account %d\n", amount, source.ID, destination.ID)
	fmt.Printf("Account %d balance: %g\n", source.ID, source.Balance)
	fmt.Printf("Account %d balance: %g\n", destination.ID, destination.Balance)
}

func main() {
	account1 := Account{ID: 1, Balance: 100.0}
	account2 := Account{ID: 2, Balance: 50.0}

	transferFunds(50.0, &account1, &account2)
	transferFunds(200.0, &account1, &account2)
}
