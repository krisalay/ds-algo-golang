package main

import "fmt"

func hasCommonItems(array1 []byte, array2 []byte) bool {
	hashTable := make(map[byte]bool)
	for _, item := range array1 {
		if !hashTable[item] {
			hashTable[item] = true
		}
	}

	for _, item := range array2 {
		if hashTable[item] {
			return true
		}
	}
	return false
}

func main() {
	array1 := []byte{'a', 'b', 'c', 'x'}
	array2 := []byte{'y', 'z', 'i'}
	result1 := hasCommonItems(array1, array2)
	fmt.Printf("Array has common items? %t \n", result1)
	array3 := []byte{'a', 'b', 'c', 'x'}
	array4 := []byte{'y', 'z', 'x'}
	result2 := hasCommonItems(array3, array4)
	fmt.Printf("Array has common items? %t", result2)
}
