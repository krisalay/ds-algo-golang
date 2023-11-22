package main

import "fmt"

func mergeSortedArrays(array1 []int, array2 []int) []int {
	if len(array1) == 0 {
		return array2
	}
	if len(array2) == 0 {
		return array1
	}

	firstPointer := 0
	secondPointer := 0

	mergedArray := []int{}

	for firstPointer < len(array1) && secondPointer < len(array2) {
		if array1[firstPointer] < array2[secondPointer] {
			mergedArray = append(mergedArray, array1[firstPointer])
			firstPointer++
		} else {
			mergedArray = append(mergedArray, array2[secondPointer])
			secondPointer++
		}
	}

	for firstPointer < len(array1) {
		mergedArray = append(mergedArray, array1[firstPointer])
		firstPointer++
	}

	for secondPointer < len(array2) {
		mergedArray = append(mergedArray, array2[secondPointer])
		secondPointer++
	}

	return mergedArray
}

func main() {
	array1 := []int{0, 3, 4, 31}
	array2 := []int{4, 6, 30}

	mergedArray := mergeSortedArrays(array1, array2)
	fmt.Println("Merged array: ", mergedArray)
}
