package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	testCase := "123"

	expected := 123
	result, err := ParseInt(testCase)

	if !assert.EqualValues(t, expected, result) {
		t.Errorf("expected value %v but got %v", expected, result)
	}

	if !assert.EqualValues(t, nil, err) {
		t.Errorf("expected error %v but got %v", nil, err)
	}
}

func TestParseFloat64Comma(t *testing.T) {
	t.Run("String satisfying the input condition ", func(t *testing.T) {
		testCase := "12,33"
		result, err := ParseFloat64Comma(testCase)
		expected := 12.33
		if !assert.EqualValues(t, expected, result) {
			t.Errorf("expected value %v but got %v", expected, result)
		}

		if !assert.EqualValues(t, nil, err) {
			t.Errorf("expected error %v but got %v", nil, err)
		}
	})

	t.Run("String does not satisfying the input condition ", func(t *testing.T) {
		testCase := "12.33"
		result, err := ParseFloat64Comma(testCase)
		expected := 12.33
		if !assert.EqualValues(t, expected, result) {
			t.Errorf("expected value %v but got %v", expected, result)
		}

		if !assert.EqualValues(t, nil, err) {
			t.Errorf("expected error %v but got %v", nil, err)
		}
	})
}

func TestParseFloat64(t *testing.T) {
	testCase := "12,33"
	result, err := ParseFloat64Comma(testCase)
	expected := 12.33
	if !assert.EqualValues(t, expected, result) {
		t.Errorf("expected value %v but got %v", expected, result)
	}

	if !assert.EqualValues(t, nil, err) {
		t.Errorf("expected error %v but got %v", nil, err)
	}
}
