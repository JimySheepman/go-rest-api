package helper

import "sort"

func FindMaxAndMinValue(arr []int) (minValue int, maxValue int) {
	sort.Ints(arr)
	return arr[0], arr[len(arr)-1]
}
