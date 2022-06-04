package homework

func reverse(input []int64) (result []int64) {
	slicecopy := []int64{}

	for i := 0; i < len(input); i++ {
		slicecopy = append(slicecopy, input[len(input)-1-i])
	}

	return slicecopy
}
