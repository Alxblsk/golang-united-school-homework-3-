package homework

func average(input [15]float32) (result float32) {
	var all float32 = 0
	for _, v := range input {
		all = all + v
	}
	return all / float32(len(input))
}
