package main

import "fmt"

func GetName() string {
	name := ""
	fmt.Printf("Enter your name: ")
	_,err := fmt.Scanln(&name)
	if err!= nil {
		return ""
	}
	fmt.Printf("Welcome %s, Let's play\n",name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet or 0 to quit (balance = $%d): ",balance)
		fmt.Scan((&bet))
		if bet > balance {
			fmt.Printf("Not enough balance (balance = $%d)\n",balance)
		} else {
			break
		}
	}
	return bet
}
