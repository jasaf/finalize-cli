package types

import "time"

type UserInput struct {
	FileName string
	FromDate time.Time
	ToDate   time.Time
}

type RowItem struct {
	Date        time.Time
	Category    string
	Subcategory string
	Amount      float64
	Currency    string
	Region      string
}

type CurrencyTotal map[string]float64

type SubCategoryWithCurrencyTotal map[string]CurrencyTotal

type CategoryTotals map[string]SubCategoryWithCurrencyTotal
