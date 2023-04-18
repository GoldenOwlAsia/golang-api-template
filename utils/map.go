package utils

func ExtractMapKeys[T, V comparable](myMap map[T]V) (result []T) {
	result = make([]T, len(myMap))
	i := 0
	for k := range myMap {
		result[i] = k
		i++
	}

	return
}

func GetCsvHeaderIndex(line []string) (output map[string]int) {
	output = make(map[string]int)
	for i, v := range line {
		output[v] = i
	}
	return
}
