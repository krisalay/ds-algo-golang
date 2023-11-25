package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		input  []int
		expect bool
	}{
		// The table itself
		{"should find duplicate in the array {1, 2, 3, 1}", []int{1, 2, 3, 1}, true},
		{"should not find duplicate in the array {1, 2, 3, 4}", []int{1, 2, 3, 4}, false},
		{"should return false for empty array", []int{}, false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := containsDuplicate(tt.input)
			if ans != tt.expect {
				t.Errorf("got %t, expect %t", ans, tt.expect)
			}
		})
	}
}
