package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Pomog/ascii-art/functions"
)

func main() {
	args := os.Args[1:]
	checkArgsLen(args)

	// Parse flags and get color value
	colorFlag := parseColorFlag()
	fmt.Println("Color:", colorFlag)

	for i, arg := range args {
		fmt.Printf("Argument %v: %s\n", i, arg)
	}

	// remove color flag from args if present
	if getArgumentsCount(args) == 4 {
		args = args[2:]
	}

	// get string from args wich will be converted to ascii art
	unquotedString, errUnquot := strconv.Unquote((`"` + args[0] + `"`))
	if errUnquot != nil {
		fmt.Println("Error Unquote:", errUnquot)
		os.Exit(1)
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	result := functions.GetProcededSclice(mapOfSymbols, unquotedString)

	functions.PrintResult(result)

	// write string ascii art to the result.txt file
	errWrite := functions.WriteToFile("result.txt", result)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	farewell("result.txt")
}

/*
print farewell message
*/
func farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!", resultFileName)
	fmt.Println(message)
}

/*
if a -color flag is present, it should be first argument folowed by a <letters to be colored> as second argument
if there is no -color flag, then the color is white
and it shoud be 2 arguments: first - string, second - name of file with symbols
exit with error if number of arguments is wrong
*/
func checkArgsLen(args []string) {
	argumentsCount := getArgumentsCount(args)

	if len(args) != argumentsCount {
		errorMessage := fmt.Sprintf(
			"Error: wrong number of arguments\nExpected: %d arguments\nGot: %d arguments\nUsage: go run main.go \"input string\" [BANNER]",
			argumentsCount, len(args))
		log.Fatal(errorMessage)
	}
}

/*
if a -color flag is present, it should be first argument folowed by a <letters to be colored> as second argument
if there is no -color flag, then the color is white by default
return color value as string
*/
func parseColorFlag() string {
	var color string
	flag.StringVar(&color, "color", "white", "Specify a color")
	flag.Parse()

	return color
}

/*
return expected number of arguments
2 if no -color flag: first - string, second - name of file with symbols ascii-art
4 if -color flag is present: first - -color flag color definition without spaces, second - <letters to be colored>

	third - string, fourth - name of file with symbols ascii-art
*/
func getArgumentsCount(args []string) int {
	if strings.Contains(args[0], "color") {
		return 4
	} else {
		return 2
	}
}
