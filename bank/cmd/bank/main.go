package main

import (
	"fmt"
	"os"

	"bank/internal/account"
	"bank/internal/utils"
)

func main() {
	acc := account.NewAccountAccess()

	var option string
	for {
		fmt.Println("Select an option\n 1. Credit\n 2. Debit\n 3. Check Balance\n 4. Exit")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Print("Enter amount: ")
			input := utils.TakeInput()

			amount, err := utils.ConvertStringAmountToInt(input)
			if err != nil {
				fmt.Println("Invalid Input")
				break
			}

			txn := acc.Credit(amount)
			if txn {
				fmt.Println("Successful")
			} else {
				fmt.Println("Failure")
			}
			break
		case "2":
			fmt.Print("Enter amount: ")
			input := utils.TakeInput()

			amount, err := utils.ConvertStringAmountToInt(input)
			if err != nil {
				fmt.Println("Invalid Input")
				break
			}

			txn := acc.Debit(amount)
			if txn {
				fmt.Println("Successful")
			} else {
				fmt.Println("Failure")
			}
			break
		case "3":
			fmt.Println("Balance: ", utils.ConvertAmountToString(acc.Balance()))
		case "4":
			os.Exit(0)
		}
	}
}
