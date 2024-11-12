package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/jasaf/finalize-cli/internal/helpers"
	"github.com/jasaf/finalize-cli/internal/types"
)

func ParseUserInput(fs *flag.FlagSet, args []string) (*types.UserInput, error) {
	var inputFlag string
	var filterDateFlag string

	fs.StringVar(&inputFlag, "input", "", "Insert CSV file name with extension")
	fs.StringVar(&filterDateFlag, "date", "", "Insert date range with the following format [from:to]")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if len(inputFlag) < 1 {
		return nil, errors.New("file name input flag is empty")
	}
	if len(filterDateFlag) < 1 {
		return nil, errors.New("filter date flag is empty")
	}

	betweenDates := strings.Split(filterDateFlag, ":")

	if len(betweenDates) < 2 {
		return nil, errors.New("could not find the dates separator")
	}

	fromDate, err := helpers.StringToDate(betweenDates[0])
	if err != nil {
		return nil, err
	}

	toDate, err := helpers.StringToDate(betweenDates[1])
	if err != nil {
		return nil, err
	}

	return &types.UserInput{
		FileName: inputFlag,
		FromDate: fromDate,
		ToDate:   toDate,
	}, nil
}

func PrintJson(data any) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}
