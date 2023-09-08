package functions

import (
	"fmt"
	"os"
	"strings"
)

/*
input: args, flag string "-flag"
*/
func IsFlagPresent(args []string, flag string) bool {
	for _, arg := range args {
		if strings.Contains(arg, flag) {
			return true
		}
	}
	return false
}

func HandleReverseFlag(reverseFlag bool, fileNameWithSymbols string, fileNameToRead string) {
	if reverseFlag {
		var reversStringFromAsciiArt string = parseAsciiArtFile(fileNameWithSymbols, fileNameToRead)
		fmt.Printf("%s\n", reversStringFromAsciiArt)
		os.Exit(0)
	}
}

/*
returns a string that is the result of reversing ASCII ART.
input: fileNameWithSymbols - the name of the file to read ASCII ART symbols representation from
input: fileNameToRead - name of the file to read ASCII ART from
ONLY ONE ASCII ART SYMBOLS LINE IS SUPPORTED
*/
func parseAsciiArtFile(fileNameWithSymbols, fileNameToRead string) string {
	mapOfSymbols, err := MakeSymbolsMapFromFile(fileNameWithSymbols)
	CheckErrorAndFatal(err)

	reversStringRecursive := GetStringFromASCIIArtRecursive(
		GetSymbolsMapVerticalRepresentation(mapOfSymbols),
		ReadFromTxtFileVertical(fileNameToRead))
	return reversStringRecursive
}
