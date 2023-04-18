package utils

import (
	"reflect"
	"testing"
)

func TestExtractMapKeys(t *testing.T) {
	testCases := []struct {
		name  string
		input map[string]int
		want  []string
	}{
		{
			name: "Test Case 1",
			input: map[string]int{
				"foo": 1,
				"bar": 2,
				"baz": 3,
			},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name:  "Test Case 2",
			input: map[string]int{},
			want:  []string{},
		},
		{
			name:  "Test Case 3",
			input: nil,
			want:  []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractMapKeys(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ExtractMapKeys(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
