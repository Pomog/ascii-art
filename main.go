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
var lettersToBeColored string = ""
var validColors = []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
var validAligns = []string{"left", "center", "right", "justify"}

func main() {
	functions.PrintTerminalWidth()

	args := os.Args[1:]

	colorFlagPresent := isColorFlagPresent(args)

	// Parse color flag as string and check if it is valid color, default color is white
	colorFlag, alignFlag := processFlags(args)
	// remove flags from args
	args = flag.Args()

	fmt.Printf("alignFlag: %s\n", alignFlag)

	if colorFlagPresent {
		lettersToBeColored = args[0]
		args = args[1:]
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(args[1] + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// get string from args wich will be converted to ascii-art, proceded string is the first element of args
	unquotedString, errUnquot := strconv.Unquote((`"` + args[0] + `"`))
	if errUnquot != nil {
		log.Fatal(errUnquot)
	}

	// obtain and combine all ascii-art symbols into the one slice of strings by layers to wokr with hole string
	result := functions.GetProcessedSlice(mapOfSymbols, unquotedString, lettersToBeColored, colorFlag)

	functions.PrintResult(result)

	// write string ascii art to the file result.txt. The color flag is not used in the file, lettersToBeColored not taken into account.
	errWrite := functions.WriteToTxtFile(resultFileName, mapOfSymbols, unquotedString)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	farewell(resultFileName)
}

/*
print farewell message
*/
func farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!", resultFileName)
	fmt.Println(message)
}

/*
isValidColor checks if a given color string is one of the valid colors.
Valid colors include: red, orange, yellow, green, blue, indigo, violet.
*/
func isValidColor(color string) bool {
	lowercaseColor := strings.ToLower(color)

	for _, validColor := range validColors {
		if lowercaseColor == validColor {
			return true
		}
	}
	return false
}

func isValidAlign(align string) bool {
	lowercaseAlign := strings.ToLower(align)

	for _, validAlign := range validAligns {
		if lowercaseAlign == validAlign {
			return true
		}
	}
	return false
}

/*
if there is no -color flag, then the color is white by default
if there is no -align flag, then the align is left by default
*/
func processFlags(args []string) (string, string) {
	var colorFlag, alignFlag string

	flag.StringVar(&colorFlag, "color", "white", "Specify a color")
	flag.StringVar(&alignFlag, "align", "left", "Specify alignment (left, center, right, justify)")
	flag.Parse()

	if !isValidColor(colorFlag) {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}

	if !isValidAlign(alignFlag) {
		alignErr := "Error: wrong align value\nExpected: one of the aligns: left, center, right, justify\nGot: " + alignFlag
		log.Fatal(alignErr)
	}

	return colorFlag, alignFlag
}

func isColorFlagPresent(args []string) bool {
	for _, arg := range args {
		if strings.Contains(arg, "-color") {
			return true
		}
	}
	return false
}
