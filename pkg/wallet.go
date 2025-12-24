package pkg

import "errors"

type Waller interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}
type BitcoinWallet struct {
	name    string
	balance int
}

type UsdWallet struct {
	name    string
	balance int
}

var (
	ErrAmountBelowZero   = errors.New("the sum must be greater than zero")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
