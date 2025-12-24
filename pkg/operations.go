package pkg

import (
	"fmt"
	"strconv"
)

func (u UsdWallet) Deposit(amount float64) {

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

	u.balance += amount

	return
}

func (b BitcoinWallet) Deposit(amount float64) {

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

	b.balance += amount

	return
}

func (b BitcoinWallet) Withdraw(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if b.balance < amount {
		fmt.Println("Error:", ErrAmountBelowZero)
		return
	}

	b.balance -= amount

}

func (u UsdWallet) Withdraw(amount float64) {

	var sAmount string
	fmt.Println("Enter the deposit amount: ")
	_, err := fmt.Scan(&sAmount)
	if err != nil {
		return
	}

	amount, err = strconv.ParseFloat(sAmount, 64)
	if u.balance < amount {
		fmt.Println("Error:", ErrAmountBelowZero)
		return
	}

	u.balance -= amount

}
