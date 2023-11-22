
# Move Zeroes

Given a string, write a function to reverse that string.

For example:

Input String = "DS Algo using Golang"
Output String = "gnaloG gnisu oglA SD"


## Solution
- swap: This function takes an array (arr) and two indices (firstIndex and secondIndex) as parameters and swaps the elements at those indices in the array.

- moveNonZerosLeft: This function takes an array (arr) as input and modifies it by moving all non-zero elements to the left while maintaining the relative order of other elements. The function uses two pointers (leftPointer and rightPointer) to traverse the array.

	- leftPointer represents the position where the next non-zero element should be placed.
	- rightPointer represents the position of the next element to be examined.
	The function iterates through the array using these pointers:

	- If the element at leftPointer is already non-zero, it increments both pointers.
	- If the element at leftPointer is zero and the element at rightPointer is non-zero, it swaps the elements at these positions and increments both pointers.
	- If the element at leftPointer is zero and the element at rightPointer is also zero, it increments only the rightPointer to find the next non-zero element.
	- This process continues until the rightPointer reaches the end of the array.

#### Code
```go
func swap(arr []int, firstIndex, secondIndex int) {
	temp := arr[firstIndex]
	arr[firstIndex] = arr[secondIndex]
	arr[secondIndex] = temp
}

func moveZeroes(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	leftPointer := 0
	rightPointer := 1

	for leftPointer < len(arr) && rightPointer < len(arr) {
		if arr[leftPointer] != 0 {
			leftPointer += 1
			rightPointer = leftPointer + 1
		} else if arr[leftPointer] == 0 && arr[rightPointer] != 0 {
			swap(arr, leftPointer, rightPointer)
			leftPointer += 1
			rightPointer = leftPointer + 1
		} else {
			rightPointer++
		}
	}
	return arr
}
```