package pkg

import "errors"

type Waller interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}
type BitcoinWallet struct {
	Name    string
	Balance float64
}

type UsdWallet struct {
	Name    string
	Balance float64
}

var (
	ErrAmountBelowZero   = errors.New("the sum must be greater than zero")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNoBalance         = errors.New("you have not created a balance")
)
