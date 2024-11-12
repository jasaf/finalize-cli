package item

import (
	"reflect"
	"testing"
	"time"

	"github.com/jasaf/finalize-cli/internal/helpers"
	"github.com/jasaf/finalize-cli/internal/types"
)

func TestCategoriseItems(t *testing.T) {
	parseDate := func(dateStr string) time.Time {
		date, err := helpers.StringToDate(dateStr)
		if err != nil {
			t.Fatalf("Failed to parse date: %v", err)
		}
		return date
	}

	tests := []struct {
		name     string
		items    []types.RowItem
		fromDate time.Time
		toDate   time.Time
		expected *types.CategoryTotals
	}{
		{
			name:     "Single item within date range",
			items:    []types.RowItem{{Date: parseDate("2024-01-01"), Category: "Food", Subcategory: "Groceries", Amount: 100, Currency: "USD"}},
			fromDate: parseDate("2024-01-01"),
			toDate:   parseDate("2024-12-31"),
			expected: &types.CategoryTotals{
				"Food": {
					"Groceries": {"USD": 100},
				},
			},
		},
		{
			name: "Multiple items across categories and subcategories",
			items: []types.RowItem{
				{Date: parseDate("2024-01-01"), Category: "Food", Subcategory: "Groceries", Amount: 100, Currency: "USD"},
				{Date: parseDate("2024-02-01"), Category: "Food", Subcategory: "Dining", Amount: 50, Currency: "USD"},
				{Date: parseDate("2024-01-15"), Category: "Transport", Subcategory: "Taxi", Amount: 20, Currency: "EUR"},
				{Date: parseDate("2024-03-01"), Category: "Food", Subcategory: "Groceries", Amount: 30, Currency: "USD"},
			},
			fromDate: parseDate("2024-01-01"),
			toDate:   parseDate("2024-12-31"),
			expected: &types.CategoryTotals{
				"Food": {
					"Groceries": {"USD": 130},
					"Dining":    {"USD": 50},
				},
				"Transport": {
					"Taxi": {"EUR": 20},
				},
			},
		},
		{
			name:     "Item outside date range",
			items:    []types.RowItem{{Date: parseDate("2023-12-31"), Category: "Food", Subcategory: "Groceries", Amount: 100, Currency: "USD"}},
			fromDate: parseDate("2024-01-01"),
			toDate:   parseDate("2024-12-31"),
			expected: &types.CategoryTotals{},
		},
		{
			name: "Multiple currencies for the same subcategory",
			items: []types.RowItem{
				{Date: parseDate("2024-01-01"), Category: "Food", Subcategory: "Groceries", Amount: 100, Currency: "USD"},
				{Date: parseDate("2024-01-02"), Category: "Food", Subcategory: "Groceries", Amount: 200, Currency: "EUR"},
			},
			fromDate: parseDate("2024-01-01"),
			toDate:   parseDate("2024-12-31"),
			expected: &types.CategoryTotals{
				"Food": {
					"Groceries": {"USD": 100, "EUR": 200},
				},
			},
		},
		{
			name:     "No items provided",
			items:    []types.RowItem{},
			fromDate: parseDate("2024-01-01"),
			toDate:   parseDate("2024-12-31"),
			expected: &types.CategoryTotals{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			got := CategoriseItems(&tt.items, tt.fromDate, tt.toDate)

			// Compare the actual and expected results
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CategoriseItems() = %v, want %v", got, tt.expected)
			}
		})
	}
}
