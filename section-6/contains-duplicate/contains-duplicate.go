package main

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
