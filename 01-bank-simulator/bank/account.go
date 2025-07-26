package bank

import "errors"

type Account struct {
	balance float64
}

func NewAccount(balance float64) *Account {
	return &Account{balance: balance}
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) float64 {
	a.balance += amount
	return  a.balance
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return errors.New("cannot withdraw, insufficient funds!")
	}

	a.balance -= amount

	return nil
}