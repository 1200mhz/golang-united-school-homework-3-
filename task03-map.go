package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) []string {
	var keys []int
	var result []string

	for k, _ := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		result = append(result, input[v])
	}

	return result
}
