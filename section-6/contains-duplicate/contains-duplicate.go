package main

import "fmt"

func containsDuplicate(arr []int) bool {
	hashMap := make(map[int]bool)

	for _, num := range arr {
		if hashMap[num] {
			return true
		}
		hashMap[num] = true
	}
	return false
}

func main() {
	result1 := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(result1)
	result2 := containsDuplicate([]int{1, 2, 3, 4})
	fmt.Println(result2)
}
