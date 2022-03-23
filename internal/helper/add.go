package helper

import (
	"strconv"
)

func Add(arr []string) int {
	result := 0
	for _, v := range arr {
		intVar, _ := strconv.Atoi(v)
		result += intVar
	}
	return result
}
