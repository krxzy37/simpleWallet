package pkg

import "errors"

type Waller interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}
type BitcoinWallet struct {
	name    string
	balance float64
}

type UsdWallet struct {
	name    string
	balance float64
}

var (
	ErrAmountBelowZero   = errors.New("the sum must be greater than zero")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNoBalance         = errors.New("you have not created a balance")
)
