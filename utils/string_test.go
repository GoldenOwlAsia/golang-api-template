package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRegNumber(t *testing.T) {
	// Test with valid registration number
	err := CheckRegNumber("DU68HU")
	assert.NoError(t, err)

	// Test with registration number length less than 5
	err = CheckRegNumber("DU68")
	assert.EqualError(t, err, "invalid registration number")

	// Test with registration number length greater than 8
	err = CheckRegNumber("ABCDEF1234")
	assert.EqualError(t, err, "invalid registration number")
}

func TestFixRegNum(t *testing.T) {
	// Test with registration number containing spaces and lowercase letters
	s := FixRegNum("du 68 hu")
	assert.Equal(t, "DU68HU", s)

	// Test with empty registration number
	s = FixRegNum("")
	assert.Empty(t, s)
}
