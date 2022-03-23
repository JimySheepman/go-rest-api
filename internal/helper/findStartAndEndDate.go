package helper

import "sort"

func FindStartAndEndDate(arr []string) (string, string) {
	sort.Strings(arr)
	return arr[0], arr[len(arr)-1]
}
