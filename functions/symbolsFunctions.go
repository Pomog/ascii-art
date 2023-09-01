package functions

import (
	"fmt"
	"log"
	"regexp"
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
func GetProcessedSlice(mapOfSymbols map[rune][]string, inputString, lettersToBeColored, colorFlag, alignFlag string) []string {
	if !CheckForNotAllowedSymbols(inputString) {
		log.Fatal(errNotAllowedSymbols) //error can be retured to the caller here and after. but I think it's better to use log.Fatal in this case
	}

	var result []string
	for _, row := range splitStringByNewline(inputString) {
		result = append(result, composeResultingSlice(mapOfSymbols, row, lettersToBeColored, colorFlag)...)
	}

	implementAlignFlag(alignFlag, &result)

	return result
}

/*
implementAlignFlag implements the align flag using pointers to the result slice.
*/
func implementAlignFlag(alignFlag string, result *[]string) {
	*result = alignText(*result, alignFlag)
}

/*
alignText aligns the input string to the left, center, right or justify of the terminal window.
*/
func alignText(result []string, alignmentFlag string) []string {
	var alignedResult []string

	for _, row := range result {
		rowLengthWithOutColor := len(regexp.MustCompile("\033\\[[0-9;]*m").ReplaceAllString(row, ""))
		padding := getAlignmentPadding(alignmentFlag, rowLengthWithOutColor)
		alignedResult = append(alignedResult, padding+row)
	}

	return alignedResult
}

/*
modified the input string to be aligned to the left, center, right or justify of the terminal window.
by adding spaces
*/
func getAlignmentPadding(alignmentFlag string, rowWidth int) string {
	terminalWidth, err := getTerminalWidth()
	CheckErrorAndFatal(err)

	switch alignmentFlag {
	case "right":
		return strings.Repeat(" ", terminalWidth-rowWidth)
	case "center":
		return strings.Repeat(" ", (terminalWidth-rowWidth)/2)
	case "justify": // TODO: implement justify
		return ""
	default:
		return ""
	}
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
	// // uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
	fd := uintptr(syscall.Stdout) //represents the file descriptor constant for the standard output stream
	ws := &winsize{}              //This instance will be used to store the retrieved terminal window size information.

	//Making a System Call (syscall.Syscall)
	/*
		syscall.SYS_IOCTL: This is a constant representing the IOCTL system call, which is used for device control operations.
		fd: The file descriptor for standard output.
		uintptr(syscall.TIOCGWINSZ): A constant representing the TIOCGWINSZ operation code, used to retrieve terminal window size information.
		uintptr(unsafe.Pointer(ws)): A pointer to the winsize struct instance, cast to uintptr using unsafe.Pointer.
	*/
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws)))
	if err != 0 {
		return 0, err
	}

	return int(ws.Col), nil
}

/*
winsize is a struct that contains the terminal width and height.
*/
type winsize struct {
	Row    uint16 //unsigned 16-bit integer
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

/*
system calls to interact with the underlying operating system to obtain information about the terminal
*/
func PrintTerminalWidth() {
	width, err := getTerminalWidth()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(width)
}

func CheckErrorAndFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetSymbolsMapVerticalRepresentation(mapOfSymbols map[rune][]string) map[rune][]string {
	var result = make(map[rune][]string)

	for symbol, symbolRepresentation := range mapOfSymbols {
		verticalRows := generateVerticalRepresentation(symbolRepresentation)
		result[symbol] = verticalRows
	}

	return result
}

func generateVerticalRepresentation(lines []string) []string {
	var verticalRow string
	var newRows []string

	for i := 0; i < len(lines[0]); i++ {
		for _, row := range lines {
			verticalRow += string(row[i])
		}
		newRows = append(newRows, verticalRow)
		verticalRow = ""
	}

	return newRows
}

func GetStringFromASCIIArt(mapOfSymbols map[rune][]string) string {
	asciiArtVetticalSlice := ReadFromTxtFileVertical("result.txt")
	var resultingString string = ""

	mapOfSymbolsVewrtical := GetSymbolsMapVerticalRepresentation(mapOfSymbols)

	for len(asciiArtVetticalSlice) > 0 {
		foundMatch := false

		for symbol, symbolRepresentation := range mapOfSymbolsVewrtical {
			symbolRepresentationLength := len(symbolRepresentation)

			if len(asciiArtVetticalSlice) >= symbolRepresentationLength && slicesAreEqual(asciiArtVetticalSlice[:symbolRepresentationLength], symbolRepresentation) {
				resultingString += string(symbol)
				asciiArtVetticalSlice = asciiArtVetticalSlice[symbolRepresentationLength:]
				foundMatch = true
				break // Exit the inner loop when a match is found
			}
		}

		if !foundMatch {
			if len(asciiArtVetticalSlice) > 0 {
				// If no match was found in this iteration and asciiArtVetticalSlice is not empty, remove the first element
				asciiArtVetticalSlice = asciiArtVetticalSlice[1:]
			} else {
				break // Exit the outer loop when asciiArtVetticalSlice is empty
			}
		}
	}

	return resultingString
}

func slicesAreEqual(slice1, slice2 []string) bool {
	// Check if slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each element of the slices
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
