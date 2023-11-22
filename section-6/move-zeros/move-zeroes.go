package main

import "fmt"

func swap(arr []int, firstIndex, secondIndex int) {
	temp := arr[firstIndex]
	arr[firstIndex] = arr[secondIndex]
	arr[secondIndex] = temp
}

func moveZeros(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	leftPointer := 0
	rightPointer := 1

	for leftPointer < len(arr) && rightPointer < len(arr) {
		if arr[leftPointer] != 0 {
			leftPointer += 1
			rightPointer = leftPointer + 1
		} else if arr[leftPointer] == 0 && arr[rightPointer] != 0 {
			swap(arr, leftPointer, rightPointer)
			leftPointer += 1
			rightPointer = leftPointer + 1
		} else {
			rightPointer++
		}
	}
	return arr
}

func main() {
	array := []int{1, 2, 3, 0, 12}
	result := moveZeros(array)
	fmt.Println(result)
}
