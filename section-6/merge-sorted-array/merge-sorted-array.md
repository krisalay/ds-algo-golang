
# Merge Two Sorted Arrays

Defines a function mergeSortedArrays that takes two sorted arrays (array1 and array2) as input and returns a new array that is a sorted combination of the input arrays

  
### Solution
- Parameters: array1 and array2 are two input arrays to be merged.
- The function first checks if either of the arrays is empty. If one of them is empty, it returns the other array as the result since merging with an empty array results in the same array.
- Two pointers (firstPointer and secondPointer) are initialized to zero, pointing to the beginning of each array.
- An empty array called mergedArray is initialized to store the merged result.
- A for loop is used to iterate through both arrays until one of them is fully processed:
- Inside the loop, it compares the elements pointed to by the two pointers (array1[firstPointer] and array2[secondPointer]).
- The smaller element is appended to the mergedArray, and the respective pointer is incremented.
- After the loop, if there are remaining elements in either array, they are appended to the mergedArray using two separate for loops.


Both the time and space complexities are O(m + n), where m and n are the lengths of the input arrays.


  

#### Code:

```go

func mergeSortedArrays(array1 []int, array2 []int) []int {
	if len(array1) == 0 {
		return array2
	}
	if len(array2) == 0 {
		return array1
	}

	firstPointer := 0
	secondPointer := 0

	mergedArray := []int{}

	for firstPointer < len(array1) && secondPointer < len(array2) {
		if array1[firstPointer] < array2[secondPointer] {
			mergedArray = append(mergedArray, array1[firstPointer])
			firstPointer++
		} else {
			mergedArray = append(mergedArray, array2[secondPointer])
			secondPointer++
		}
	}

	for firstPointer < len(array1) {
		mergedArray = append(mergedArray, array1[firstPointer])
		firstPointer++
	}

	for secondPointer < len(array2) {
		mergedArray = append(mergedArray, array2[secondPointer])
		secondPointer++
	}

	return mergedArray
}
```