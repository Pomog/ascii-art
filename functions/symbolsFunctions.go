package functions

import (
	"fmt"
	"os"
)

var errNotAllowedSymbols = "Not allowed symbols in input string"

func GetProcededSclice(mapOfSymbols map[rune][]string, inputString string) []string {
	if checkForNotAllowedSymbols(inputString) {
		fmt.Println(errNotAllowedSymbols)
		os.Exit(1)
	}

	var result []string
	for _, row := range newLineSybolHandling(inputString) {
		result = append(result, composeResultingSlice(mapOfSymbols, row)...)
	}
	return result
}

func checkForNotAllowedSymbols(inputString string) bool {
	for _, symbol := range inputString {
		if (symbol <= 32 && symbol >= 126) || symbol != 10 {
			return false
		}
	}
	return true
}

func newLineSybolHandling(inputString string) []string {
	var slicedString []string
	var row string
	for _, symbol := range inputString {
		if symbol == '\n' {
			slicedString = append(slicedString, row)
			row = ""
		} else {
			row += string(symbol)
		}
	}

	if row != "" {
		slicedString = append(slicedString, row)
	}

	return slicedString
}

func composeResultingSlice(mapOfSymbols map[rune][]string, row string) []string {
	var result []string
	var currentRow string

	if row != "" {
		for i := 0; i < symbolHeight; i++ {
			for _, symbol := range row {
				currentRow += mapOfSymbols[symbol][i]
			}
			result = append(result, currentRow)
			currentRow = ""
		}
	} else {
		result = append(result, "")
	}
	return result
}

func PrintResult(result []string) {
	for _, row := range result {
		fmt.Println(row)
	}
}
