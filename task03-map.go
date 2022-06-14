package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	myKeys := []int{}
	for k := range input {
		myKeys = append(myKeys, k)
	}
	sort.Ints(myKeys)
	myValues := []string{}
	for _, k := range myKeys {
		myValues = append(myValues, input[k])
	}
	return myValues
}
