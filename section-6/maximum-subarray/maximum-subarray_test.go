package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		name   string
		input  []int
		expect int
	}{
		// The table itself
		{"should return 6 for array {-2, 1, -3, 4, -1, 2, 1, -5, 4}", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := maxSubArray(tt.input)
			if ans != tt.expect {
				t.Errorf("got %d, expect %d", ans, tt.expect)
			}
		})
	}
}
