package utils

import (
	"strconv"
	"strings"
)

func ParseInt(value string) (result int, err error) {
	result, err = strconv.Atoi(value)
	return
}

func ParseFloat64(value string) (result float64, err error) {
	result, err = strconv.ParseFloat(value, 64)
	return
}

func ParseFloat64Comma(value string) (result float64, err error) {
	value = strings.Replace(value, ",", ".", 1)
	result, err = strconv.ParseFloat(value, 64)
	return
}
