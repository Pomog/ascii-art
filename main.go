package main

import (
	"fmt"
	"os"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	result := functions.GetProcededSclice(mapOfSymbols, "{Hello+\nThere}")
	for _, row := range result {
		fmt.Println(row)
	}
}
