# Contains Duplicate
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

## Solution
### Approach #1 (Naive Linear Search) [Time Limit Exceeded]

#### Algorithm

To apply this idea, we employ the linear search algorithm which is the simplest search algorithm. Linear search is a method of finding if a particular value is in a list by checking each of its elements, one at a time and in sequence until the desired one is found.

For our problem, we loop through all n integers. For the ith integer nums[i], we search in the previous i-1 integers for the duplicate of nums[i]. If we find one, we return true; if not, we continue. Return false at the end of the program.

To prove the correctness of the algorithm, we define the loop invariant. A loop invariant is a property that holds before (and after) each iteration. Knowing its invariant(s) is essential for understanding the effect of a loop. Here is the loop invariant:

Before the next search, there are no duplicate integers in the searched integers.

The loop invariant holds true before the loop because there is no searched integer.
Each time through the loop we look for any any possible duplicate of the current element.
If we found a duplicate, the function exits by returning true; If not, the invariant still holds true.

Therefore, if the loop finishes, the invariant tells us that there is no duplicate in all nn integers.



### Approach #2 (Sorting)
#### Intuition

If there are any duplicate integers, they will be consecutive after sorting.

#### Algorithm

This approach employs sorting algorithm. Since comparison sorting algorithm like heapsort is known to provide O(nlog⁡n) worst-case performance, sorting is often a good preprocessing step. After sorting, we can sweep the sorted array to find if there are any two consecutive duplicate elements.

#### Note

The implementation here modifies the original array by sorting it. In general, it is not a good practice to modify the input unless it is clear to the caller that the input will be modified. One may make a copy of nums and operate on the copy instead.

### Approach #3 (Hash Table)
#### Intuition

Utilize a dynamic data structure that supports fast search and insert operations.

#### Algorithm

From Approach #1 we know that search operations is O(n) in an unsorted array and we did so repeatedly. Utilizing a data structure with faster search time will speed up the entire algorithm.

There are many data structures commonly used as dynamic sets such as Binary Search Tree and Hash Table. The operations we need to support here are search() and insert(). For a self-balancing Binary Search Tree, search() and insert() are both O(log⁡n) time. For a Hash Table (HashSet or HashMap), search() and insert() are both O(1) on average. Therefore, by using hash table, we can achieve linear time complexity for finding the duplicate in an unsorted array.

#### Code
```go
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
```

#### Complexity Analysis

Time complexity: O(n).
We do search() and insert() for n times and each operation takes constant time.

Space complexity: O(n).
The space used by a hash table is linear with the number of elements in it.

#### Note

For certain test cases with not very large n, the runtime of this method can be slower than Approach #2. The reason is hash table has some overhead in maintaining its property. One should keep in mind that real world performance can be different from what the Big-O notation says. The Big-O notation only tells us that for sufficiently large input, one will be faster than the other. Therefore, when n is not sufficiently large, an O(n) algorithm can be slower than an O(nlog⁡n) algorithm.
