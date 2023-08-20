package functions

import (
	"bufio"
	"os"
)

var strtRune = 32
var symbolHeight = 8

func ReadFromFile(fileName string) (map[rune][]string, error) {
	file, errRead := os.Open(fileName)
	if errRead != nil {
		return nil, errRead
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	return processSymbols(fileScanner)
}

func processSymbols(fileScanner *bufio.Scanner) (map[rune][]string, error) {
	symbols := make(map[rune][]string)
	var currentSymbolName rune
	var currentSymbol []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			currentSymbol = append(currentSymbol, line)
			if len(currentSymbol) == symbolHeight {
				currentSymbolName = rune(strtRune)
				symbols[currentSymbolName] = currentSymbol
				strtRune++
				currentSymbol = nil
			}
		}
	}

	if errScan := fileScanner.Err(); errScan != nil {
		return nil, errScan
	}

	return symbols, nil
}
