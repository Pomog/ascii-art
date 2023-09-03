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
var fileNameWithSymbolsDefault = "standard.txt"
var lettersToBeColored string = ""
var validColors = []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet", "white"}
var validAligns = []string{"left", "center", "right", "justify"}
var validAsciiArtSourse = []string{"standard", "shadow", "thinkertoy"}

func main() {
	args := os.Args[1:]

	// this var needed to help parse lettersToBeColored
	colorFlagPresent := isFlagPresent(args, "-color")

	// this var needed to handle case when -reverse flag is present, program will print reversed string and exit
	reverseFlagPresent := isFlagPresent(args, "-reverse")

	// Parse color and align flag as strings and check if they are valid
	colorFlag, alignFlag, reverseFlag := processFlags(args)
	// remove flags from args
	args = flag.Args()

	// if -reverse flag is present, then reverse the string print it and exit
	if reverseFlagPresent {
		var reversStringFromAsciiArt string = parseAsciiArtFile(fileNameWithSymbolsDefault, reverseFlag)

		fmt.Printf("%s\n", reversStringFromAsciiArt)
		os.Exit(0)
	}

	// parsing and removing lettersToBeColored from args
	if colorFlagPresent && len(args) == 3 {
		lettersToBeColored = args[0]
		args = args[1:]
	} else if !colorFlagPresent && len(args) == 3 {
		log.Fatal("Error: wrong number of arguments\nFlag -color is not present but lettersToBeColored included")
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols := getMapOfSymbols(args, validAsciiArtSourse)

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
func processFlags(args []string) (string, string, string) {
	colorFlag, alignFlag, reverseFlag := parseFlags()

	if !isValueValid(colorFlag, validColors) {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}

	if !isValueValid(alignFlag, validAligns) {
		alignErr := "Error: wrong align value\nExpected: one of the aligns: left, center, right, justify\nGot: " + alignFlag
		log.Fatal(alignErr)
	}

	if reverseFlag != "" && !fileIsPresent(reverseFlag) {
		reverseErr := "Error: File " + reverseFlag + " is not found\n"
		log.Fatal(reverseErr)
	}

	return colorFlag, alignFlag, reverseFlag
}

/*
parseFlags parses the flags and returns the color and align flags as strings.
*/
func parseFlags() (string, string, string) {
	var colorFlag, alignFlag, reverseFlag string

	flag.StringVar(&colorFlag, "color", "white", "Specify a color")
	flag.StringVar(&alignFlag, "align", "left", "Specify alignment (left, center, right, justify)")
	flag.StringVar(&reverseFlag, "reverse", "", "Reverse the string")
	flag.Parse()
	return colorFlag, alignFlag, reverseFlag
}

/*
getMapOfSymbols returns a map of symbols from the file.
The usage must respect this format go run . [STRING] [BANNER]
*/
func getMapOfSymbols(args []string, validAsciiArtSourse []string) map[rune][]string {
	if isValueValid(args[len(args)-1], validAsciiArtSourse) {
		mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
		functions.CheckErrorAndFatal(err)
		return mapOfSymbols
	} else {
		log.Fatal("Usage: go run . [STRING] [BANNER]\nEX: go run . \"something\" standard")
		return nil
	}
}

/*
returns a string that is the result of reversing ASCII ART.
input: fileNameWithSymbols - the name of the file to read ASCII ART symbols representation from
input: fileNameToRead - name of the file to read ASCII ART from
ONLY ONE ASCII ART SYMBOLS LINE IS SUPPORTED
*/
func parseAsciiArtFile(fileNameWithSymbols, fileNameToRead string) string {
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(fileNameWithSymbols)
	functions.CheckErrorAndFatal(err)

	reversStringRecursive := functions.GetStringFromASCIIArtRecursive(
		functions.GetSymbolsMapVerticalRepresentation(mapOfSymbols),
		functions.ReadFromTxtFileVertical(fileNameToRead))
	return reversStringRecursive
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

func fileIsPresent(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
input: args, flag string "-flag"
*/
func isFlagPresent(args []string, flag string) bool {
	for _, arg := range args {
		if strings.Contains(arg, flag) {
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
