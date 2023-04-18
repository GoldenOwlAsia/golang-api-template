package utils

import "testing"

func TestHumanFileSize(t *testing.T) {
	testCases := []struct {
		input  float64
		output string
	}{
		{100, "100 B"},
		{1024, "1 KB"},
		{1024 * 1024, "1 MB"},
		{1024 * 1024 * 1024, "1 GB"},
		{1024 * 1024 * 1024 * 1024, "1 TB"},
	}

	for _, tc := range testCases {
		result := HumanFileSize(tc.input)
		if result != tc.output {
			t.Errorf("Expected %v but got %v", tc.output, result)
		}
	}
}
