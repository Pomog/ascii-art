package helpers

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pomog/ascii-art/configuration"
	"github.com/Pomog/ascii-art/functions"
)

// Access the configuration using config.ConfigInstance
var config = configuration.ConfigInstance

/*
if there is no -color flag, then the color is white by default
if there is no -align flag, then the align is left by default
*/
func ProcessFlags(args []string) (string, string, string, error) {
	colorFlag, alignFlag, reverseFlag := ParseFlags()

	if !IsValueValid(colorFlag, config.ValidColors) {
		colorErr := "Error: wrong color value\nExpected: one of the colors: red, orange, yellow, green, blue, indigo, violet\nGot: " + colorFlag
		log.Fatal(colorErr)
	}

	if !IsValueValid(alignFlag, config.ValidAligns) {
		alignErr := "Error: wrong align value\nExpected: one of the aligns: left, center, right, justify\nGot: " + alignFlag
		log.Fatal(alignErr)
	}

	if reverseFlag != "" && !FileIsPresent(reverseFlag) {
		reverseErr := "Error: File " + reverseFlag + " is not found\n"
		log.Fatal(reverseErr)
	}

	return colorFlag, alignFlag, reverseFlag, nil
}

/*
parseFlags parses the flags and returns the color and align flags as strings.
*/
func ParseFlags() (string, string, string) {
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
func GetMapOfSymbols(banner string) map[rune][]string {
	mapOfSymbols, err := functions.MakeSymbolsMapFromFile(banner + ".txt")
	functions.CheckErrorAndFatal(err)
	return mapOfSymbols
}

/*
isValueValid checks if the value is in the slice of allowed values.
*/
func IsValueValid(value string, allowedValues []string) bool {
	for _, allowedValue := range allowedValues {
		if strings.ToLower(value) == allowedValue {
			return true
		}
	}
	return false
}

func FileIsPresent(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
print farewell message
*/
func Farewell(resultFileName string) {
	message := fmt.Sprintf("Finished. No errors. Thanks for using.\nThe result is in the file --> %s <--\nGoodbye!", resultFileName)
	fmt.Println(message)
}

// parseArguments parses and validates command-line arguments and updates the config struct.
func ParseArguments(args []string, config *configuration.Config) error {
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
		if !IsValueValid(config.Banner, config.ValidAsciiArtSourse) {
			return fmt.Errorf("error: wrong banner value\nExpected: one of the banners: %s\nGot: %s", config.ValidAsciiArtSourse, config.Banner)
		}
	case 2:
		if IsValueValid(args[1], config.ValidAsciiArtSourse) {
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
