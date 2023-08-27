package functions

import (
	"bufio"
	"os"
)

var strtRune = 32
var symbolHeight = 8

func MakeSymbolsMapFromFile(fileName string) (map[rune][]string, error) {
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

func WriteToTxtFile(fileName string, mapOfSymbols map[rune][]string, inputString string) error { // TODO: unit tests
	result := GetProcessedSlice(mapOfSymbols, inputString, "", "white")

	file, errWrite := os.Create(fileName)
	if errWrite != nil {
		return errWrite
	}
	defer file.Close()

	fileWriter := bufio.NewWriter(file)
	for _, line := range result {
		fileWriter.WriteString(line + "\n")
	}

	// Flush any buffered data to the file
	if errFlush := fileWriter.Flush(); errFlush != nil {
		return errFlush
	}
	file.Close() //deferred closure may not be executed if an error occurs before the defer statement
	return nil
}
