package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	numbs := make([]int, len(input))
	res := make([]string, len(input))

	var counter = 0
	for i := range input {
		numbs[counter] = i
		counter++
	}

	sort.Ints(numbs)

	for i, v := range numbs {
		res[i] = input[v]
	}

	return res
}
