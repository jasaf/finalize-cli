package main

import (
	"log"

	"github.com/jasaf/finalize-cli/internal/console"
	"github.com/jasaf/finalize-cli/internal/file"
	"github.com/jasaf/finalize-cli/internal/item"
)

func main() {
	userInput, err := console.ParseUserInput()
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
