package functions

import (
	"fmt"
	"os"
)

var errNotAllowedSymbols = "Not allowed symbols in input string"

func GetProcededSclice(mapOfSymbols map[rune][]string, inputString string) []string {
	if checkforAllowedSymbols(inputString) {
		fmt.Println(errNotAllowedSymbols + ": " + inputString)
		os.Exit(1)
	}

	var result []string
	for _, row := range newLineSybolHandling(inputString) {
		result = append(result, composeResultingSlice(mapOfSymbols, row)...)
	}
	return result
}

func checkforAllowedSymbols(inputString string) bool {
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
		if symbol == 10 {
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
	var finalRow string
	for i := 0; i < symbolHeight; i++ {
		for _, symbol := range row {
			finalRow += mapOfSymbols[symbol][i]
		}
		result = append(result, finalRow)
		finalRow = ""
	}
	return result
}
