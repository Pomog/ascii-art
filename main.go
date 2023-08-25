package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	args := os.Args[1:]

	// Parse flags and get color value
	parseFlags()

	for i, arg := range args {
		fmt.Printf("Argument %v: %s\n", i, arg)
	}

	checkArgs(args) // it shoud be 2 arguments: first - string, second - name of file with symbols

	if len(args) == 4 {
		// remove first 2 arguments
		args = append(args[:0], args[2:]...)
	}

	unquotedString, errUnquot := strconv.Unquote((`"` + args[0] + `"`))
	if errUnquot != nil {
		fmt.Println("Error Unquote:", errUnquot)
		os.Exit(1)
	}

	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
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

func checkArgs(args []string) {
	argumentsCount := 2

	if strings.Contains(args[0], "color") {
		argumentsCount = 4
	}

	if len(args) != argumentsCount {
		fmt.Println("Error: wrong number of arguments")
		fmt.Println("Expected:", argumentsCount, "arguments")
		fmt.Println("Got:", len(args), "arguments")
		fmt.Println("Usage: go run main.go \"input string\" [BANNER]")
		os.Exit(1)
	}
}

func parseFlags() string {
	colorPtr := flag.String("color", "white", "Specify a color")
	flag.Parse()

	color := *colorPtr

	switch color {
	case "red":
		fmt.Println("You chose red!")
	case "green":
		fmt.Println("You chose green!")
	case "blue":
		fmt.Println("You chose blue!")
	default:
		fmt.Println("Unrecognized color:", color)
	}

	return color
}
