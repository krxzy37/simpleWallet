package main

import (
	"fmt"
	"os"
	"simpleWallet/pkg"

	"github.com/zetamatta/go-getch"
)

func main() {

	var usdWallName string
	var bitWallName string

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
			fmt.Println("Creating a wallet ....\nchoose the wallet b-> bitcoin wallet, u-> usd wallet")

			var chose string
			_, err := fmt.Scan(&chose)

			if err != nil {
				fmt.Println("write 'u' or 'b' ")
				return
			}

			if chose == "u" {
				fmt.Println("Write name of usd wallet")
				_, err := fmt.Scan(&usdWallName)
				if err != nil {
				}

				usdWallet := pkg.UsdWallet{
					Name: usdWallName,
				}

				fmt.Println("Your wallet was created", usdWallet)

			} else if chose == "b" {
				fmt.Println("Write name of usd wallet")
				_, err := fmt.Scan(&usdWallName)
				if err != nil {
				}

				bitWallName := pkg.BitcoinWallet{
					Name: bitWallName,
				}

				fmt.Println("Your wallet was created", bitWallName)

			}

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
