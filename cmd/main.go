package main

import (
	"flag"
	"log"
	"os"

	"github.com/jasaf/finalize-cli/internal/console"
	"github.com/jasaf/finalize-cli/internal/file"
	"github.com/jasaf/finalize-cli/internal/item"
)

func main() {
	fs := flag.CommandLine
	userInput, err := console.ParseUserInput(fs, os.Args[1:])
	if err != nil {
		log.Fatalf("Error while parsing user input, %s\n", err)
	}

	items, err := file.ReadCsvFile(userInput.FileName)
	if err != nil {
		log.Fatalf("Error while reading csv file, %s\n", err)
	}

	categorisedItems := item.CategoriseItems(items, userInput.FromDate, userInput.ToDate)

	if err := console.PrintJson(categorisedItems); err != nil {
		log.Fatalf("Error while printing json")
	}
}
