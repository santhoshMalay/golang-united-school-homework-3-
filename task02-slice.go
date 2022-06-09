package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		temp := input[i]
		input[i] = input[j]
		input[j] = temp
	}
	return input
}
