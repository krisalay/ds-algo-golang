# Google Software Engineer Solve an Engineering Problem
This repository provides an in-depth exploration of the problem-solving thought process, unraveling solutions to entirely unfamiliar challenges. For a detailed explanation of a Google engineer's problem-solving approach, refer to this insightful [video](https://youtu.be/XKu_SEDAykw).

## Table of contents

* [Problem Statement](#problem-statement)
* [Jot Down Details](#jot-down-details)
* [Brute Force Approach](#brute-force-approach)
* [Binary Search Approach](#binary-search-approach)
* [Two Pointer Approach](#two-pointer-approach)
* [Array is Not Sorted Anymore](#array-is-not-sorted-anymore)
* [Array is Large](#array-is-large)

### Problem Statement
Identify pairs within a collection where the sum equals a specified target.
[1, 2, 3, 9] sum = 8
[1, 2, 4, 4] sum = 8

### Jot Down Details
1. How are these collections given - in memory or array or something else? ***Response:** Array and assume that the array is sorted in ascending order*
2. What about the repeated element in the array? Can we use the same index twice? ***Response:** Same index cannot be used twice, but same number can appear at different indices. For instance 4 and 4 at indices 2 and 3 in the second example provided.* 
3. How about the elements of the array, are they integers or floating point? ***Response:** The elements are integers.*
4. Can the numbers be negatives? ***Response:** Negatives could happen.*

### Brute Force Approach
The most straightforward approach involves comparing every conceivable pair, achieved by implementing a nested for-loop. The outer loop starts from the 0th index, while the inner loop begins from the 1st index, systematically checking each pair to determine if their sum equals the target value. While not the most efficient method, this serves as one viable solution to the problem.
### Binary Search Approach
Iterating through each element in the array, we check if its complement exists. For example, in the first array, the complement of 1 is 8-1, i.e., 7. Leveraging the array's sorted nature, we employ binary search to ascertain the existence of 7. If found, we've identified a pair. This method, with a binary search time complexity of O(logn) and an O(n) for-loop, offers an improvement over quadratic approaches, resulting in an O(nlogn) complexity. While efficient, it's worth noting that this approach may be comparatively slower.

***Hint:** Instead of doing binary search, which is uni-directional, what if we start with pair of numbers to begin with and then work our way inward from there.*

### Two Pointer Approach
Assuming the left pointer starts at the 0th index and the right pointer at length-1, we evaluate the sum of their corresponding values. If the sum exceeds the target value, we shift the right pointer leftward by one index. Conversely, if the sum is less than the target, we move the left pointer rightward by one index. This process continues until the sum matches the target or the two pointers converge at a single index. This linear approach ensures an efficient solution with a time complexity of O(n).
#### Code:
```go
func  hasPairWithSumInSortedArray(arr []int, target int) bool {
	left, right  :=  0, len(arr)-1
	for left < right {
		sum  := arr[left] + arr[right]
		if sum == target {
			return  true
		} else  if sum > target {
			right--
		} else {
			left++
		}
	} 
	// If no pair is found
	return  false
}
```

> Let's make this a bit more difficult. How will you solve this same problem if that array is not sorted?

### Array is Not Sorted Anymore
The first thing ofcourse we can think is first sort the array, then we can solve the problem with the same way as above, but it still will be an O(nlogn) solution. But we want faster than that.

Instead of inquiring, 'Is there anywhere?' we shift our perspective to, 'Have I seen it in the past?' For instance, traversing the first array at index 2 with a value of 3, the complement is 8-3, i.e., 5. The crucial step is to check if we've encountered 5 before. To efficiently track this, we employ a data structure suitable for lookups, such as a hash set, ensuring constant time lookup. We insert a 7 when encountering a 1 and a 6 when seeing a 2—effectively storing the complements. The careful handling of repeated elements is essential; checking before insertion ensures a seamless process.

Time Complexity: O(n)
Space Complexity: O(n)

#### Code:
```go
func  hasPairWithSum(arr []int, target int) bool {
	compliments  :=  make(map[int]bool)
	for  _, num  :=  range arr {
		compliment  := target - num
		if compliments[compliment] {
			return  true
		}
		compliments[compliment] =  true
	}
	return  false
}
```

### Array is Large
What would we do differently if say we have 10 million integers in the array?
**Questions**
1. Does the array fit in the memory? ***Response:** No, the array does not fit in the memory.*

If our hash-map fits in memory, but input does not fit in memory, then we can sort of process it in chunks.