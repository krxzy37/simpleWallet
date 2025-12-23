package main

import (
	"fmt"
	"os"

	"github.com/zetamatta/go-getch"
	_ "github.com/zetamatta/go-getch"
)

func main() {

	fmt.Println("LocalWallet by krxzy37")

	for {
		fmt.Println("---Menu---")
		fmt.Println("1.Create Wallet")
		fmt.Println("2.top up your wallet")
		fmt.Println("3.withdraw money from a wallet")
		fmt.Println("0.Exit")

		fmt.Println()

		r := getch.Rune()

		switch r {
		case '1':
		case '2':
		case '3':
		case '0':
			fmt.Println("Exiting the program....")
			os.Exit(0)
		default:
			fmt.Println("Please enter a number")

		}
	}

}
