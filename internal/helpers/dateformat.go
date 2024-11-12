package helpers

import (
	"time"
)

const (
	defaultFormat = "2006-01-02"
)

func StringToDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse(defaultFormat, date)

	return parsedDate, err
}
