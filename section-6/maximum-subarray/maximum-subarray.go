package main

import (
	"math"
)

// Brute force approach
func maxSubarrayUsingBruteForce(arr []int) int {
	maxSubarray := math.Inf(-1)

	for i := 0; i < len(arr); i++ {
		currentSubarray := 0
		for j := i; j < len(arr); j++ {
			currentSubarray += arr[j]
			maxSubarray = math.Max(maxSubarray, float64(currentSubarray))
		}
	}
	return int(maxSubarray)
}

// Optimised Approach
func maxSubArray(arr []int) int {
	maxSubArray := arr[0]
	currentSubarray := arr[0]

	for i := 1; i < len(arr); i++ {
		num := arr[i]
		currentSubarray = int(math.Max(float64(num), float64(currentSubarray+num)))
		maxSubArray = int(math.Max(float64(currentSubarray), float64(maxSubArray)))
	}
	return maxSubArray
}

// func main() {
// 	result1 := maxSubarrayUsingBruteForce([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
// 	fmt.Println(result1)
// 	result2 := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
// 	fmt.Println(result2)
// }
