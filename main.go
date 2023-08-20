package main

import (
	"fmt"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	mapOfSymbols, err := functions.ReadFromFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for symbol, lines := range mapOfSymbols {
		fmt.Println("Symbol:", string(symbol))
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println("-----")
	}

	for _, row := range mapOfSymbols['A'] {
		fmt.Println(row)
	}
}
