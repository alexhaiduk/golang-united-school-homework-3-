package homework

func reverse(input []int64) (result []int64) {
	var startindex int = len(input) - 1
	//result = input[0:len(input)]
	//var value = input[0:len(input)]
	for i := startindex; i >= 0; i-- {
		result[startindex-i] = input[i]
	}
	return result
}
