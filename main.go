package main

import "fmt"


func main()  {
	symbols := map[string]uint {
		"A":4,
		"B":7,
		"C":12,
		"D":20,
	}

	multiplier := map[string]uint {
		"A":20,
		"B":10,
		"C":5,
		"D":2,
	}


	fmt.Println("Welcome to Casino")
	balance := uint(500)
	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}		
		balance -= bet
		symbolArr := GenerateSymbolArray(symbols)
		spin := GetSpin(symbolArr,3,3)
		PrintSpin(spin)
		winningLines := CheckWin(spin,multiplier)
		for i,multi := range winningLines {
			win := bet * multi
			balance += win
			if multi > 0 {
				fmt.Printf("Won $%d (%dX) on Line #%d\n",win,multi,i)
			}
		}
	}
	fmt.Printf("You left with $%d",balance)
}