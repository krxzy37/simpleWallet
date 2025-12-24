package pkg

import (
	"fmt"
	"strconv"
)

func (u *UsdWallet) Deposit(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if err != nil {
		fmt.Println("Type a number! Error name:", err)
		return
	}
	if amount < 0 {
		fmt.Println("Error:", ErrAmountBelowZero)
		return
	}

	u.Balance += amount

	return
}

func (b *BitcoinWallet) Deposit(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if err != nil {
		fmt.Println("Type a number! Error name:", err)
		return
	}
	if amount < 0 {
		fmt.Println("Error:", ErrAmountBelowZero)
		return
	}

	b.Balance += amount

	return
}

func (b *BitcoinWallet) Withdraw(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if b.Balance < amount {
		fmt.Println("Error:", ErrInsufficientFunds)
		return
	}

	b.Balance -= amount

}

func (u *UsdWallet) Withdraw(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if u.Balance < amount {
		fmt.Println("Error:", ErrInsufficientFunds)
		return
	}

	u.Balance -= amount

}

func (u *UsdWallet) GetBalance() float64 {

	if u == nil {
		fmt.Println("Error: ", ErrNoBalance)
		return 0.0
	}

	return u.Balance
}

func (b *BitcoinWallet) GetBalance() float64 {

	if b == nil {
		fmt.Println("Error: ", ErrNoBalance)
		return 0.0
	}

	return b.Balance
}
