package file

import (
	"bytes"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/jasaf/finalize-cli/internal/helpers"
	"github.com/jasaf/finalize-cli/internal/types"
)

func ReadCsvFile(name string) (*[]types.RowItem, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var items []types.RowItem
	for _, record := range records[1:] {
		amount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}

		date, err := helpers.StringToDate(record[0])
		if err != nil {
			return nil, err
		}

		newItem := types.RowItem{
			Date:        date,
			Category:    record[1],
			Subcategory: record[2],
			Amount:      amount,
			Currency:    record[4],
			Region:      record[5],
		}

		items = append(items, newItem)
	}

	return &items, nil
}
