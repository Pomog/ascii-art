package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	args := os.Args[1:]

	unquotedString, errUnquot := strconv.Unquote((`"` + args[0] + `"`))
	if errUnquot != nil {
		fmt.Println("Error Unquote:", errUnquot)
		os.Exit(1)
	}

	mapOfSymbols, err := functions.MakeSymbolsMapFromFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	result := functions.GetProcededSclice(mapOfSymbols, unquotedString)

	functions.PrintResult(result)

	errWrite := functions.WriteToFile("result.txt", result)
	if errWrite != nil {
		fmt.Println("Error:", errWrite)
		os.Exit(1)
	}

	farewell("result.txt")
}

func farewell(resultFileName string) {
	fmt.Println("Finished. No errors. Thanks for using.")
	fmt.Printf("The result is in the file --> %s <--\n", resultFileName)
	fmt.Println("Goodbye!")
}
