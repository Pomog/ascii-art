package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pomog/ascii-art/functions"
)

// Initialize the config struct
var config = initConfig()

func main() {
	args := os.Args[1:]

	// this var needed to handle case when -reverse flag is present, program will print reversed string and exit
	reverseFlagPresent := functions.IsFlagPresent(args, "-reverse")

	// Parse color and align flag as strings and check if they are valid
	colorFlag, alignFlag, reverseFlag, errFlagsProcessing := processFlags(args)
	if errFlagsProcessing != nil {
		log.Fatal(errFlagsProcessing)
	}
	// remove flags from args
	args = flag.Args()

	//if -reverse flag is present, then reverse the string print it and exit
	functions.HandleReverseFlag(reverseFlagPresent, config.FileNameWithSymbolsDefault, reverseFlag)

	// parse and validate command-line arguments, after all flags are processed
	err := parseArguments(args, &config)
	if err != nil {
		log.Fatal(err)
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols := getMapOfSymbols(config.Banner)

	// get string from args wich will be converted to ascii-art, proceded string is the first element of args
	unquotedString := strings.ReplaceAll(config.InputString, "\\n", "\n")

	// obtain and combine all ascii-art symbols into the one slice of strings by layers to wokr with hole string
	result := functions.GetProcessedSlice(mapOfSymbols, unquotedString, config.LettersToBeColored, colorFlag, alignFlag)

	if alignFlag == "justify" {
		justifiedString := functions.Justify(unquotedString, mapOfSymbols)
		result := functions.GetProcessedSlice(mapOfSymbols, justifiedString, config.LettersToBeColored, colorFlag, "left")
		functions.PrintResult(result)
		os.Exit(0)
	}

	functions.PrintResult(result)

	// write string ascii art to the file result.txt. The color flag is not used in the file, lettersToBeColored not taken into account.
	errWrite := functions.WriteToTxtFile(config.ResultsFileName, mapOfSymbols, unquotedString)
	functions.CheckErrorAndFatal(errWrite)

	farewell(config.ResultsFileName)
}

//	>>>    Helper functions    <<<

/*
if there is no -color flag, then the color is white by default
if there is no -align flag, then the align is left by default
*/
func processFlags(args []string) (string, string, string, error) {
	colorFlag, alignFlag, reverseFlag := parseFlags()

	if !isValueValid(colorFlag, config.ValidColors) {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}

	if !isValueValid(alignFlag, config.ValidAligns) {
		alignErr := "Error: wrong align value\nExpected: one of the aligns: left, center, right, justify\nGot: " + alignFlag
		log.Fatal(alignErr)
	}

	if reverseFlag != "" && !fileIsPresent(reverseFlag) {
		reverseErr := "Error: File " + reverseFlag + " is not found\n"
		log.Fatal(reverseErr)
	}

	return colorFlag, alignFlag, reverseFlag, nil
}

/*
parseFlags parses the flags and returns the color and align flags as strings.
*/
func parseFlags() (string, string, string) {
	var colorFlag, alignFlag, reverseFlag string

	flag.StringVar(&colorFlag, "color", "white", "Specify a color")
	flag.StringVar(&alignFlag, "align", "left", "Specify alignment (left, center, right, justify)")
	flag.StringVar(&reverseFlag, "reverse", "", "Reverse the ASCII ART from the file")

	// Parse the command-line flags
	flag.Parse()

	return colorFlag, alignFlag, reverseFlag
}

/*
getMapOfSymbols returns a map of symbols from the file.
The usage must respect this format go run . [STRING] [BANNER]
*/
func getMapOfSymbols(banner string) map[rune][]string {
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(banner + ".txt")
	functions.CheckErrorAndFatal(err)
	return mapOfSymbols
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
print farewell message
*/
func farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!", resultFileName)
	fmt.Println(message)
}

// parseArguments parses and validates command-line arguments and updates the config struct.
func parseArguments(args []string, config *Config) error {
	colorFlagPresent := functions.IsFlagPresent(args, "-color")

	if len(args) > 3 || len(args) == 0 {
		return fmt.Errorf("error: wrong number of arguments\nExpected: 1, 2 or 3 arguments\nGot: %v", len(args))
	}

	if !colorFlagPresent && len(args) == 3 {
		return fmt.Errorf("error: wrong number of arguments\nFlag -color is not present but lettersToBeColored included")
	}

	switch len(args) {
	case 3:
		config.Banner = args[2]
		config.InputString = args[1]
		config.LettersToBeColored = args[0]
		if !isValueValid(config.Banner, config.ValidAsciiArtSourse) {
			return fmt.Errorf("error: wrong banner value\nExpected: one of the banners: %s\nGot: %s", config.ValidAsciiArtSourse, config.Banner)
		}
	case 2:
		if isValueValid(args[1], config.ValidAsciiArtSourse) {
			config.Banner = args[1]
			config.InputString = args[0]
		} else {
			config.InputString = args[1]
			config.LettersToBeColored = args[0]
		}
	case 1:
		config.InputString = args[0]
	}

	return nil
}
