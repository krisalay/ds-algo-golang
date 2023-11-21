package main

import "fmt"

// This function works only with sorted array in ascending order
func hasPairWithSumInSortedArray(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return true
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	// If no pair is found
	return false
}

// This function works for any array of integers
func hasPairWithSum(arr []int, target int) bool {
	compliments := make(map[int]bool)
	for _, num := range arr {
		compliment := target - num

		if compliments[compliment] {
			return true
		}
		compliments[compliment] = true
	}
	return false
}

func main() {
	arr1 := []int{1, 2, 3, 9}
	arr2 := []int{1, 2, 4, 4}
	arr3 := []int{1, 4, 8, 9}
	target := 8
	hasPair1 := hasPairWithSumInSortedArray(arr1, target)
	fmt.Printf("Does pair exist in [1, 2, 3, 9] whose sum is 8? %t \n", hasPair1)
	hasPair2 := hasPairWithSumInSortedArray(arr2, target)
	fmt.Printf("Does pair exist in [1, 2, 4, 4] whose sum is 8? %t\n", hasPair2)
	hasPair3 := hasPairWithSum(arr3, target)
	fmt.Printf("Does pair exist in [1, 2, 4, 4] whose sum is 8? %t\n", hasPair3)
}
