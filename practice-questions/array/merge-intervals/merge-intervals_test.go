package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		// The table itself
		{"[[1,3],[2,6],[8,10],[15,18]]", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"[[1,4],[4,5]]", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"[[1,4],[2,3]]", [][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{"[[1,4],[1,5]]", [][]int{{1, 4}, {1, 5}}, [][]int{{1, 5}}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := merge(tt.input)
			if len(ans) != len(tt.expect) {
				t.Errorf("got %d, expect %d", ans, tt.expect)
			} else {
				for i, num := range ans {
					if num[0] != tt.expect[i][0] || num[1] != tt.expect[i][1] {
						t.Errorf("got %d, expect %d", ans, tt.expect)
					}
				}
			}
		})
	}
}
