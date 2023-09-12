package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Pomog/ascii-art/configuration"
	"github.com/Pomog/ascii-art/functions"
	"github.com/Pomog/ascii-art/helpers"
	"github.com/gin-gonic/gin"
)

var AsciiArt []string
var config = configuration.ConfigInstance

func main() {
	args := os.Args[1:]

	// this var needed to handle case when -reverse flag is present, program will print reversed string and exit
	reverseFlagPresent := functions.IsFlagPresent(args, "-reverse")

	// Parse flags as strings and check if they are valid
	colorFlag, alignFlag, reverseFlag, errFlagsProcessing := helpers.ProcessFlags(args)
	if errFlagsProcessing != nil {
		log.Fatal(errFlagsProcessing)
	}
	// remove flags from args
	args = flag.Args()

	//if -reverse flag is present, then print resulting string it and exit
	functions.HandleReverseFlag(reverseFlagPresent, config.FileNameWithSymbolsDefault, reverseFlag)

	// parse and validate command-line arguments, after all flags are processed, results stored in config struct
	err := helpers.ParseArguments(args, config)
	if err != nil {
		log.Fatal(err)
	}

	// get map of symbols from file, where key is a symbol and value is a slice of strings wich represents the symbol
	mapOfSymbols := helpers.GetMapOfSymbols(config.Banner)

	// get string from args wich will be converted to ascii-art, proceded string is the first element of args
	unquotedString := strings.ReplaceAll(config.InputString, "\\n", "\n")

	// obtain and combine all ascii-art symbols into the one slice of strings by layers to wokr with hole string
	result := functions.GetProcessedSlice(mapOfSymbols, unquotedString, config.LettersToBeColored, colorFlag, alignFlag)

	/*
		TODO: need to be implemented using consistent logic
	*/
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

	/*
		Server part using gin framework
	*/
	AsciiArt = functions.GetASCIIARTSlice(unquotedString, mapOfSymbols)

	router := gin.Default()
	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/asciiart", RenderStringsPage)

	router.Run("localhost:9090")

	helpers.Farewell(config.ResultsFileName)
}
