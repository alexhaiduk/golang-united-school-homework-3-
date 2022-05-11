package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input)+1)
	var l int = 1
	for i := 0; i < len(input); i++ {
		result[i] = input[len(input)-l]
		l++
	}
	return
}
