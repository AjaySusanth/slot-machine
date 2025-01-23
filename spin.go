package main

import (
	"fmt"
	"math/rand"
)


func GetRandomNumber (min int, max int) int {
	randomNumber := rand.Intn(max - min + 1) + min
	return randomNumber
}

func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol,count := range symbols {
		for i:=uint(0); i<count;i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func GetSpin (reel []string, rows int,cols int) [][]string {
	result := [][]string{}
	for i:=0; i<rows;i++ {
		result = append(result, []string{})
	}
	for col:=0;col<cols;col++ {
		selected:= map[int]bool{}

		for row:=0 ;row<rows;row++ {
			for true {
				randIndex := GetRandomNumber(0,len(reel)-1)
				_,exists :=selected[randIndex]

				if !exists {
					selected[randIndex] = true
					result[row] = append(result[row],reel[randIndex])
					break
				}
			}
		}

	}
	return result
}

func PrintSpin(spin [][]string) {
	for _,row := range spin {
		for j,symbol := range row {
			fmt.Print(symbol)
			if j != len(row) -1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func CheckWin (spin [][]string, multiplier map[string]uint) []uint {
	lines := []uint {}
	for _,row := range spin {
		win :=true
		checkSymbol:= row[0]
		for _,symbol := range row[1:]{
			if checkSymbol != symbol {
				win = false
			}
		}
		if win {
			lines = append(lines, multiplier[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}