package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	val := map[int]string{
		9:  "Сентябрь",
		1:  "Январь",
		2:  "Февраль",
		10: "Октябрь",
		5:  "Май",
		7:  "Июль",
		8:  "Август",
		12: "Декарь",
		3:  "Март",
		6:  "Июнь",
		4:  "Апрель",
		11: "Ноябрь",
	}

	GetMapValuesSortedByKey(val)
}

// ReturnInt - Function returns int value
func ReturnInt() int {
	return 1
}

// ReturnFloat - Function returns float value
func ReturnFloat() float32 {
	return 1.1
}

// ReturnIntArray - Function returns array of int values
func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

// ReturnIntSlice - Function returns slice of int values
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

// IntSliceToString - Function convert int slice to string
func IntSliceToString(elements []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(elements)), ""), "[]")
}

// MergeSlices - Merge two different slices
func MergeSlices(slice1 []float32, slice2 []int32) []int {
	slice := make([]int, len(slice1)+len(slice2))
	for key, value := range slice1 {
		slice[key] = int(value)
	}
	for key, value := range slice2 {
		slice[key+len(slice1)] = int(value)
	}
	return slice
}

// GetMapValuesSortedByKey - Sort map
func GetMapValuesSortedByKey(m map[int]string) []string {
	var keys []int
	var results []string

	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		results = append(results, m[key])
	}

	return results
}
