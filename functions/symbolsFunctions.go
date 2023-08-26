package functions

import (
	"fmt"
	"log"
	"strings"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Orange = "\033[38;5;208m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Indigo = "\033[38;5;63m"
	Violet = "\033[38;5;129m"
)

var ErrNotAllowedSymbols = "Not allowed symbols in the input string"

func GetProcededSclice(mapOfSymbols map[rune][]string, inputString, lettersToBeColored, colorFlag string) []string {
	if !CheckForNotAllowedSymbols(inputString) {
		log.Fatal(ErrNotAllowedSymbols)
	}

	var result []string
	for _, row := range newLineSybolHandling(inputString) {
		result = append(result, composeResultingSlice(mapOfSymbols, row, lettersToBeColored, colorFlag)...)
	}
	return result
}

func CheckForNotAllowedSymbols(inputString string) bool {
	for _, symbol := range inputString {
		if symbol < 32 || symbol > 127 {
			fmt.Println(string(symbol))
			fmt.Println(symbol)
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

func composeResultingSlice(mapOfSymbols map[rune][]string, row, lettersToBeColored, colorFlag string) []string {
	var result []string
	var currentRow string

	if row != "" {
		for i := 0; i < symbolHeight; i++ {
			for _, symbol := range row {
				currentRow += processSymbol(symbol, mapOfSymbols, lettersToBeColored, colorFlag, i)
			}
			result = append(result, currentRow)
			currentRow = ""
		}
	} else {
		result = append(result, "")
	}
	return result
}

func processSymbol(symbol rune, mapOfSymbols map[rune][]string, lettersToBeColored, colorFlag string, i int) string {
	if lettersToBeColored != "" && strings.Contains(lettersToBeColored, string(symbol)) {
		return colorize(mapOfSymbols[symbol][i], colorFlag)
	}
	return mapOfSymbols[symbol][i]
}

func colorize(row, colorFlag string) string {
	switch colorFlag {
	case "red":
		return Red + row + Reset
	case "orange":
		return Orange + row + Reset
	case "yellow":
		return Yellow + row + Reset
	case "green":
		return Green + row + Reset
	case "blue":
		return Blue + row + Reset
	case "indigo":
		return Indigo + row + Reset
	case "violet":
		return Violet + row + Reset
	default:
		return row
	}
}

func PrintResult(result []string) {
	for _, row := range result {
		fmt.Println(row)
	}
}
