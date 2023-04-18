package utils

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	dateStr := "2023-03-20"
	expectedDate := time.Date(2023, 03, 20, 0, 0, 0, 0, time.UTC)

	result, err := ParseDate(dateStr)

	if err != nil {
		t.Errorf("Unexpected error while parsing date: %v", err)
	}

	if result != expectedDate {
		t.Errorf("Unexpected date returned. Expected %v, but got %v", expectedDate, result)
	}
}

func TestParseDateTime(t *testing.T) {
	// Test cases with valid input strings
	testCases := []struct {
		input    string
		expected time.Time
	}{
		{"2022-03-20T12:34:56Z", time.Date(2022, time.March, 20, 12, 34, 56, 0, time.UTC)},
		{"2022-01-01T00:00:00-08:00", time.Date(2022, time.January, 1, 0, 0, 0, 0, time.FixedZone("", -8*60*60))},
	}

	for _, tc := range testCases {
		actual, err := ParseDateTime(tc.input)
		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", tc.input, err)
			continue
		}
		if !actual.Equal(tc.expected) {
			t.Errorf("Unexpected result for input %q. Expected %v, but got %v", tc.input, tc.expected, actual)
		}
	}

	// Test cases with invalid input strings
	invalidInputs := []string{"", "2022-03-20", "not a valid datetime string"}
	for _, input := range invalidInputs {
		_, err := ParseDateTime(input)
		if err == nil {
			t.Errorf("Expected an error for input %q, but got none", input)
		}
	}
}

func TestDatetimeToString(t *testing.T) {
	// create a test time
	testTime := time.Date(2023, time.March, 20, 10, 30, 0, 0, time.UTC)
	// expected result
	expectedResult := "2023-03-20T10:30:00Z"

	// call the function
	result, err := DatetimeToString(testTime)

	// check for errors
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// compare the result with the expected result
	if result != expectedResult {
		t.Errorf("expected %q but got %q", expectedResult, result)
	}
}
