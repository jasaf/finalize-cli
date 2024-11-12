package helpers

import (
	"testing"
	"time"
)

func TestStringToDate(t *testing.T) {
	tests := []struct {
		name        string    // Name of the test case
		input       string    // Input date string
		expected    time.Time // Expected parsed date
		expectError bool      // Whether an error is expected
	}{
		{
			name:        "Valid date format",
			input:       "2024-01-01",
			expected:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Invalid date format (wrong delimiter)",
			input:       "2024/01/01",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid date format (missing month)",
			input:       "2024-01",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Empty date string",
			input:       "",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid date value",
			input:       "2024-13-01", // Month out of range
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			got, err := StringToDate(tt.input)

			// Check if an error was expected
			if (err != nil) != tt.expectError {
				t.Fatalf("StringToDate(%q) error = %v, expectError %v", tt.input, err, tt.expectError)
			}

			// Compare the parsed date with the expected date
			if !got.Equal(tt.expected) {
				t.Errorf("StringToDate(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
