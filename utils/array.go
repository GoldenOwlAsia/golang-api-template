package utils

func ChunkSlice[T comparable](slice []T, chunkSize int) (chunks [][]T) {
	if len(slice) == 0 {
		return [][]T{}
	}
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func Intersect(a, b []string) []string {
	m := make(map[string]bool)
	for _, x := range a {
		m[x] = true
	}
	var res []string
	for _, x := range b {
		if m[x] {
			res = append(res, x)
		}
	}
	return res
}

func Unique[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
