package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pomog/ascii-art/functions"
)

var resultsFileName = "result.txt"
var lettersToBeColored string = ""
var validColors = []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet", "white"}
var validAligns = []string{"left", "center", "right", "justify"}

func main() {
	args := os.Args[1:]

	colorFlagPresent := isColorFlagPresent(args)

	// Parse color and align flag as strings and check if they are valid
	colorFlag, alignFlag := processFlags(args)
	// remove flags from args
	args = flag.Args()

	// parsing and removing lettersToBeColored from args
	if colorFlagPresent {
		lettersToBeColored = args[0]
		args = args[1:]
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
	functions.CheckErrorAndFatal(err)

	// reversString := functions.GetStringFromASCIIArt(mapOfSymbols)
	// fmt.Printf("\nreversString: %s\n", reversString)

	// reversStringRecursive := functions.GetStringFromASCIIArtRecursive(functions.GetSymbolsMapVerticalRepresentation(mapOfSymbols), functions.ReadFromTxtFileVertical("result.txt"))
	// fmt.Printf("\nreversStringRecursive: %s\n", reversStringRecursive)

	// get string from args wich will be converted to ascii-art, proceded string is the first element of args
	unquotedString := strings.ReplaceAll(args[0], "\\n", "\n")

	// obtain and combine all ascii-art symbols into the one slice of strings by layers to wokr with hole string
	result := functions.GetProcessedSlice(mapOfSymbols, unquotedString, lettersToBeColored, colorFlag, alignFlag)

	functions.PrintResult(result)

	// write string ascii art to the file result.txt. The color flag is not used in the file, lettersToBeColored not taken into account.
	errWrite := functions.WriteToTxtFile(resultsFileName, mapOfSymbols, unquotedString)
	functions.CheckErrorAndFatal(errWrite)

	farewell(resultsFileName)
}

//	>>>    Helper functions    <<<

/*
if there is no -color flag, then the color is white by default
if there is no -align flag, then the align is left by default
*/
func processFlags(args []string) (string, string) {
	colorFlag, alignFlag := parseFlags()

	if !isValueValid(colorFlag, validColors) {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}

	if !isValueValid(alignFlag, validAligns) {
		alignErr := "Error: wrong align value\nExpected: one of the aligns: left, center, right, justify\nGot: " + alignFlag
		log.Fatal(alignErr)
	}

	return colorFlag, alignFlag
}

/*
parseFlags parses the flags and returns the color and align flags as strings.
*/
func parseFlags() (string, string) {
	var colorFlag, alignFlag string

	flag.StringVar(&colorFlag, "color", "white", "Specify a color")
	flag.StringVar(&alignFlag, "align", "left", "Specify alignment (left, center, right, justify)")
	flag.Parse()
	return colorFlag, alignFlag
}

/*
isValueValid checks if the value is in the slice of allowed values.
*/
func isValueValid(value string, allowedValues []string) bool {
	for _, allowedValue := range allowedValues {
		if strings.ToLower(value) == allowedValue {
			return true
		}
	}
	return false
}

func isColorFlagPresent(args []string) bool {
	for _, arg := range args {
		if strings.Contains(arg, "-color") {
			return true
		}
	}
	return false
}

/*
print farewell message
*/
func farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!", resultFileName)
	fmt.Println(message)
}
