package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var i int = 0
	var sorted = make([]int, len(input))
	for k, _ := range input {
		sorted[i] = k
		i++
	}
	sort.Ints(sorted)
	result = make([]string, len(sorted))
	i = 0
	for k := range sorted {
		result[i] = input[sorted[k]]
		i++
	}
	return
}
