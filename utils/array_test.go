package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkSlice(t *testing.T) {
	t.Run("Chunk size less than slice length", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
		chunkSize := 3

		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8},
		}

		result := ChunkSlice(slice, chunkSize)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v but got %v", expected, result)
		}
	})

	t.Run("Chunk size greater than slice length", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		chunkSize := 5

		expected := [][]string{
			{"apple", "banana", "cherry"},
		}

		result := ChunkSlice(slice, chunkSize)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v but got %v", expected, result)
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		slice := []int{}
		chunkSize := 5

		expected := [][]int{}

		result := ChunkSlice(slice, chunkSize)

		if !assert.EqualValues(t, expected, result) {
			t.Errorf("expected %v but got %v", expected, result)
		}
	})
}

func TestIntersect(t *testing.T) {
	a := []string{"apple", "banana", "orange", "pear"}
	b := []string{"orange", "pear", "kiwi", "pineapple"}

	expected := []string{"orange", "pear"}
	actual := Intersect(a, b)

	if len(actual) != len(expected) {
		t.Errorf("Intersect(%v, %v) = %v; expected %v", a, b, actual, expected)
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Intersect(%v, %v) = %v; expected %v", a, b, actual, expected)
		}
	}
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, []int{1, 2, 3, 4}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		result := Unique(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Unique(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
