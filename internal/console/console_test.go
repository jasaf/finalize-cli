package console

import (
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/jasaf/finalize-cli/internal/helpers"
	"github.com/jasaf/finalize-cli/internal/types"
)

func TestParseUserInput(t *testing.T) {
	parseDate := func(dateStr string) time.Time {
		t.Helper()

		date, err := helpers.StringToDate(dateStr)
		if err != nil {
			t.Fatalf("Failed to parse date %s: %v", dateStr, err)
		}
		return date
	}

	tests := []struct {
		name        string           // Name of the test case
		args        []string         // Command-line arguments
		want        *types.UserInput // Expected output
		expectError bool             // Whether an error is expected
	}{
		{
			name: "Valid input flags",
			args: []string{"-input", "input.csv", "-date", "2024-01-01:2024-02-01"},
			want: &types.UserInput{
				FileName: "input.csv",
				FromDate: parseDate("2024-01-01"),
				ToDate:   parseDate("2024-02-01"),
			},
			expectError: false,
		},
		{
			name:        "Missing input flag",
			args:        []string{"-date", "2024-01-01:2024-02-01"},
			want:        nil,
			expectError: true,
		},
		{
			name:        "Missing date flag",
			args:        []string{"-input", "input.csv"},
			want:        nil,
			expectError: true,
		},
		{
			name:        "Invalid date format",
			args:        []string{"-input", "input.csv", "-date", "2024-01-01"},
			want:        nil,
			expectError: true,
		},
		{
			name:        "Invalid date range separator",
			args:        []string{"-input", "input.csv", "-date", "2024-01-01/2024-02-01"},
			want:        nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new FlagSet for each test case
			fs := flag.NewFlagSet(tt.name, flag.ExitOnError)

			// Call ParseUserInput with the custom FlagSet and arguments
			got, err := ParseUserInput(fs, tt.args)

			if (err != nil) != tt.expectError {
				t.Fatalf("ParseUserInput() error = %v, expectError %v", err, tt.expectError)
			}

			// Compare the actual and expected results
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUserInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
