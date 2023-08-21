package main

import (
	"fmt"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := functions.GetProcededSclice(mapOfSymbols, "Hello\nWorld")
	for _, row := range result {
		fmt.Println(row)
	}
}
