// Problem Statement
// https://leetcode.com/problems/merge-intervals

package main

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := [][]int{}
	placeholder := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= placeholder[1] {
			placeholder[1] = int(math.Max(float64(placeholder[1]), float64(intervals[i][1])))
		} else {
			result = append(result, placeholder)
			placeholder = intervals[i]
		}
	}
	result = append(result, placeholder)
	return result
}
