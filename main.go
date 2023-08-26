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

var resultFileName = "result.txt"
var lettersToBeColored string

func main() {
	args := os.Args[1:]
	checkArgsLen(args)

	// Parse color flag as string and check if it is valid color, default color is white
	colorFlag := processColorFlag(args)

	// remove color flag from args if present
	if getArgumentsCount(args) == 4 {
		lettersToBeColored = args[1]
		args = args[2:]
	}

	// get string from args wich will be converted to ascii-art, proceded string is the first element of args
	unquotedString, errUnquot := strconv.Unquote((`"` + args[0] + `"`))
	if errUnquot != nil {
		log.Fatal(errUnquot)
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// obtain and combine all ascii-art symbols into the one slice of strings by layers to wokr with hole string
	result := functions.GetProcededSclice(mapOfSymbols, unquotedString, lettersToBeColored, colorFlag)

	functions.PrintResult(result)

	// write string ascii art to the file
	errWrite := functions.WriteToFile(resultFileName, result)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	farewell(resultFileName)
}

/*
processColorFlag checks if a given color string is one of the valid colors.
Valid colors include: red, orange, yellow, green, blue, indigo, violet.
*/
func processColorFlag(args []string) string {
	colorFlag := parseColorFlag()
	if !isValidColor(colorFlag) && getArgumentsCount(args) == 4 {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}
	return colorFlag
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
	var colorStr string
	flag.StringVar(&colorStr, "color", "white", "Specify a color")
	flag.Parse()

	return colorStr
}

/*
return expected number of arguments
2 if no -color flag: first - string, second - name of file with symbols ascii-art
4 if -color flag is present:
** first - color flag - color definition without spaces,
** second - <letters to be colored>
** third - string,
** fourth - name of file with symbols ascii-art
*/
func getArgumentsCount(args []string) int {
	if strings.Contains(args[0], "color") {
		return 4
	} else {
		return 2
	}
}

/*
isValidColor checks if a given color string is one of the valid colors.
Valid colors include: red, orange, yellow, green, blue, indigo, violet.
*/
func isValidColor(color string) bool {
	validColors := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	lowercaseColor := strings.ToLower(color)

	for _, validColor := range validColors {
		if lowercaseColor == validColor {
			return true
		}
	}
	return false
}
