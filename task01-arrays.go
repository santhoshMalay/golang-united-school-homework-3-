package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0.0
	for _, val := range input {
		sum = sum + val
	}
	return sum / float32(len(input))
}
