package utils

import (
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"time"
)

// ParseDate value: "2023-03-20"
func ParseDate(value string) (result time.Time, err error) {
	result, err = time.Parse("2006-01-02", value)

	return
}

// ParseDate value: "20/03/2023"
func ParseDateCSV(value string) (result time.Time, err error) {
	result, err = time.Parse(configs.DefaultDateCsvLayoutFormat, value)

	return
}

// ParseDateTime value: "2022-03-20T12:34:56Z"
func ParseDateTime(value string) (result time.Time, err error) {
	result, err = time.Parse(time.RFC3339, value)

	return
}

// DatetimeToString output: "2022-03-20T12:34:56Z"
func DatetimeToString(value time.Time) (result string, err error) {
	result = value.Format(time.RFC3339)

	return
}
