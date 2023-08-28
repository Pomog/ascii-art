package functions

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"unsafe"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Orange = "\033[38;5;208m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Indigo = "\033[38;5;63m"
	Violet = "\033[38;5;129m"
)

var errNotAllowedSymbols = "Not allowed symbols in the input string"

/*
GetProcessedSlice returns a slice of strings which represents the input string as ASCII art.
The resulting slice is obtained by combining all ASCII art symbols into one slice of strings by layers to work with the whole string.
*/
func GetProcessedSlice(mapOfSymbols map[rune][]string, inputString, lettersToBeColored, colorFlag string) []string {
	if !CheckForNotAllowedSymbols(inputString) {
		log.Fatal(errNotAllowedSymbols) //error can be retured to the caller here and after. but I think it's better to use log.Fatal in this case
	}

	var result []string
	for _, row := range splitStringByNewline(inputString) {
		result = append(result, composeResultingSlice(mapOfSymbols, row, lettersToBeColored, colorFlag)...)
	}
	return result
}

/*
CheckForNotAllowedSymbols checks if the input string contains only allowed symbols.
Allowed symbols are: ASCII symbols from 32 to 127 and the newline symbol.
*/
func CheckForNotAllowedSymbols(inputString string) bool {
	for _, symbol := range inputString {
		if (symbol < 32 || symbol > 127) && symbol != 10 {
			return false
		}
	}
	return true
}

/*
splitStringByNewline splits the input string by the newline symbol and returns a slice of strings.
If the input string does not contain the newline symbol, the string is added to the slicedString
and a new string (to add to the slicedString) begins to be formed.
*/
func splitStringByNewline(inputString string) []string {
	var slicedString []string
	var row string
	for _, symbol := range inputString {
		row, slicedString = updateSlicedString(row, symbol, slicedString)
	}

	if row != "" {
		slicedString = append(slicedString, row)
	}

	return slicedString
}

/*
updateSlicedString updates the slicedString based on the current symbol.
*/
func updateSlicedString(row string, symbol rune, slicedString []string) (string, []string) {
	if symbol == '\n' {
		slicedString = append(slicedString, row)
		row = ""
	} else {
		row += string(symbol)
	}
	return row, slicedString
}

/*
composeResultingSlice returns a slice of strings which represents the input string as ASCII art.
*/
func composeResultingSlice(mapOfSymbols map[rune][]string, row, lettersToBeColored, colorFlag string) []string {
	var result []string
	var currentRow string

	if row != "" {
		for i := 0; i < symbolHeight; i++ { // symbolHeight is a global variable from fileHandler.go
			// Loop through each row of the symbol's ASCII art representation.
			// The loop variable i corresponds to the current row being processed.
			for _, symbol := range row {
				currentRow += processSymbolsRow(symbol, mapOfSymbols, lettersToBeColored, colorFlag, i)
			}
			result = append(result, currentRow)
			currentRow = ""
		}
	} else {
		result = append(result, "") // If the row is empty, add an empty string to the result.
	}
	return result
}

/*
processSymbolsRow returns a string which represents a single row of the symbol.
*/
func processSymbolsRow(symbol rune, mapOfSymbols map[rune][]string, lettersToBeColored, colorFlag string, i int) string {
	if lettersToBeColored != "" && strings.Contains(lettersToBeColored, string(symbol)) {
		return colorize(mapOfSymbols[symbol][i], colorFlag)
	}
	return mapOfSymbols[symbol][i]
}

/*
colorize returns a string which represents a single row of the symbol with color based on colorFlag
by adding ANSI color codes to the row.
*/
func colorize(row, colorFlag string) string {
	switch colorFlag {
	case "red":
		return Red + row + Reset
	case "orange":
		return Orange + row + Reset
	case "yellow":
		return Yellow + row + Reset
	case "green":
		return Green + row + Reset
	case "blue":
		return Blue + row + Reset
	case "indigo":
		return Indigo + row + Reset
	case "violet":
		return Violet + row + Reset
	default:
		return row
	}
}

/*
PrintResult prints the result slice line by line.
*/
func PrintResult(result []string) {
	for _, row := range result {
		fmt.Println(row)
	}
}

func getTerminalWidth() (int, error) {
	fd := uintptr(syscall.Stdout)
	ws := &winsize{}

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws)))
	if err != 0 {
		return 0, err
	}

	return int(ws.Col), nil
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func PrintTemunalWidth() {
	width, err := getTerminalWidth()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(width)
}
