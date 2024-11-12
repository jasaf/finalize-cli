package item

import (
	"time"

	"github.com/jasaf/finalize-cli/internal/types"
)

func CategoriseItems(items *[]types.RowItem, fromDate, toDate time.Time) *types.CategoryTotals {
	filteredItems := make(types.CategoryTotals)

	for _, record := range *items {
		if record.Date.Before(fromDate) || record.Date.After(toDate) {
			continue
		}

		_, categoryExists := filteredItems[record.Category]
		if !categoryExists {
			filteredItems[record.Category] = make(types.SubCategoryWithCurrencyTotal)
		}

		categoryMap := filteredItems[record.Category]
		_, subCategoryExists := categoryMap[record.Subcategory]
		if !subCategoryExists {
			categoryMap[record.Subcategory] = make(types.CurrencyTotal)
		}

		currencyTotal := categoryMap[record.Subcategory]
		_, currencyExists := currencyTotal[record.Currency]
		if !currencyExists {
			currencyTotal[record.Currency] = 0
		}

		currencyTotal[record.Currency] += record.Amount
	}

	return &filteredItems
}
