package main

import "fmt"

// This function reverses a string in O(n) time, but its also take O(n) space.
func reverseString(str string) string {
	if len(str) <= 1 {
		return str
	}
	backwardArray := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		backwardArray = append(backwardArray, str[i])
	}
	return string(backwardArray)
}

func main() {
	result := reverseString("DS Algo using Golang")
	fmt.Printf("Reverse of the string 'DS Algo using Golang' is '%s'", result)
}
