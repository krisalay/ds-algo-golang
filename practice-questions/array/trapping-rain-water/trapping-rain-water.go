// Problem Statement
// https://leetcode.com/problems/trapping-rain-water/description/

// For better understanding of this logic, refer to https://www.youtube.com/watch?v=C8UjlJZsHBw

package main

import (
	"fmt"
	"math"
)

// TC - O(n^2)
// SC - O(1)
func trapUsingBruteForce(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	trappedWater := 0
	width := 1
	for i := 0; i < len(height); i++ {
		leftWallHeight := 0
		rightWallHeight := 0
		for j := i; j >= 0; j-- {
			leftWallHeight = int(math.Max(float64(height[j]), float64(leftWallHeight)))
		}
		for k := i; k < len(height); k++ {
			rightWallHeight = int(math.Max(float64(height[k]), float64(rightWallHeight)))
		}
		waterLevel := int(math.Min(float64(leftWallHeight), float64(rightWallHeight)))
		trappedWater += ((waterLevel - height[i]) * width)
	}
	return trappedWater
}

type BoundaryHeight struct {
	left  int
	right int
}

// TC - O(n)
// SC - O(n)
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	width := 1
	boundaryHeights := make([]BoundaryHeight, len(height))
	for i := range boundaryHeights {
		boundaryHeights[i].left, boundaryHeights[i].right = 0, 0
	}
	for i := 1; i < len(height); i++ {
		boundaryHeights[i].left = int(math.Max(float64(height[i-1]), float64(boundaryHeights[i-1].left)))
	}

	for i := len(height) - 2; i >= 0; i-- {
		boundaryHeights[i].right = int(math.Max(float64(height[i+1]), float64(boundaryHeights[i+1].right)))
	}

	trappedWater := 0

	for i := 1; i < len(height)-1; i++ {
		waterLevel := int(math.Min(float64(boundaryHeights[i].left), float64(boundaryHeights[i].right)))
		level := waterLevel - height[i]
		if level >= 0 {
			trappedWater += (level * width)
		}
	}

	return trappedWater
}

// Can we optimize it more?
// TC - O(n)
// SC - O(1)
func trapUsingTwoPointer(height []int) int {
	totalBlocks := len(height)
	if totalBlocks <= 2 {
		return 0
	}
	trappedWater := 0
	leftBoundary, rightBoundary := height[0], height[totalBlocks-1]
	left, right := 0, totalBlocks-2

	for left <= right {
		if leftBoundary < rightBoundary {
			if height[left] >= leftBoundary {
				leftBoundary = height[left]
			} else {
				trappedWater += leftBoundary - height[left]
			}
			left += 1
		} else {
			if height[right] >= rightBoundary {
				rightBoundary = height[right]
			} else {
				trappedWater += rightBoundary - height[right]
			}
			right -= 1
		}
	}
	return trappedWater
}

func main() {
	water := trapUsingTwoPointer([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println("Trapped Water: ", water)
	water1 := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println("Trapped Water: ", water1)
	water2 := trapUsingBruteForce([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println("Trapped Water: ", water2)
}
