package file

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/jasaf/finalize-cli/internal/helpers"
	"github.com/jasaf/finalize-cli/internal/types"
)

func TestReadCsvFile(t *testing.T) {
	// Helper function to create a temporary CSV file for testing
	createTempCsv := func(t *testing.T, content string) string {
		t.Helper()

		// Create a temporary file
		file, err := os.CreateTemp("", "testdata_*.csv")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}

		// Write the content to the file
		if _, err := file.Write([]byte(content)); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		// Close the file
		if err := file.Close(); err != nil {
			t.Fatalf("Failed to close temporary file: %v", err)
		}

		// Return the file name
		return file.Name()
	}

	// Helper function to clean up temporary files
	cleanupTempCsv := func(t *testing.T, fileName string) {
		t.Helper()

		if err := os.Remove(fileName); err != nil {
			t.Logf("Failed to remove temporary file: %v", err)
		}
	}

	parseDate := func(dateStr string) time.Time {
		t.Helper()

		date, err := helpers.StringToDate(dateStr)
		if err != nil {
			t.Fatalf("Failed to parse date %s: %v", dateStr, err)
		}
		return date
	}

	// Define the test cases
	tests := []struct {
		name        string
		csvContent  string
		expected    *[]types.RowItem
		expectError bool
	}{
		{
			name: "Valid CSV file",
			csvContent: `Date,Category,Subcategory,Amount,Currency,Region
2024-01-01,Food,Groceries,100.50,USD,US
2024-01-02,Transport,Taxi,20.00,USD,US`,
			expected: &[]types.RowItem{
				{
					Date:        parseDate("2024-01-01"),
					Category:    "Food",
					Subcategory: "Groceries",
					Amount:      100.50,
					Currency:    "USD",
					Region:      "US",
				},
				{
					Date:        parseDate("2024-01-02"),
					Category:    "Transport",
					Subcategory: "Taxi",
					Amount:      20.00,
					Currency:    "USD",
					Region:      "US",
				},
			},
			expectError: false,
		},
		{
			name:        "Missing file",
			csvContent:  "",
			expected:    nil,
			expectError: true,
		},
		{
			name: "Invalid date format",
			csvContent: `Date,Category,Subcategory,Amount,Currency,Region
01-2024-01,Food,Groceries,100.50,USD,US`,
			expected:    nil,
			expectError: true,
		},
		{
			name: "Invalid amount format",
			csvContent: `Date,Category,Subcategory,Amount,Currency,Region
2024-01-01,Food,Groceries,invalid_amount,USD,US`,
			expected:    nil,
			expectError: true,
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// If the test case has CSV content, create a temporary CSV file
			var fileName string
			if tt.csvContent != "" {
				fileName = createTempCsv(t, tt.csvContent)
				defer cleanupTempCsv(t, fileName)
			} else {
				fileName = "non_existent_file.csv"
			}

			// Call the function under test
			got, err := ReadCsvFile(fileName)

			// Check if an error was expected
			if (err != nil) != tt.expectError {
				t.Fatalf("ReadCsvFile() error = %v, expectError %v", err, tt.expectError)
			}

			// Compare the actual and expected results
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ReadCsvFile() = %v, want %v", got, tt.expected)
			}
		})
	}
}
