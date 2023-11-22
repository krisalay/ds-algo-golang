# Maximum Sub Array
Given an integer array nums, find the subarray with the largest sum, and return its sum.

**Example 1:**

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]

Output: 6

Explanation: The subarray [4,-1,2,1] has the largest sum 6.

**Example 2:**

Input: nums = [1]

Output: 1

Explanation: The subarray [1] has the largest sum 1.

**Example 3:**

Input: nums = [5,4,-1,7,8]

Output: 23

Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


## Brute Force Approach

- Initialize `maxSubarray = -infinity` to keep the track of sum of best subarray. We are using negative infinity instead of zero because there is a possibility that the array many contain only negative values.
- Use a `for loop` that considers each index of the array as starting point.
- For each starting point, create a varibale `currentSubarray = 0`. Then loop through array from starting index, adding each element to `currentSubarray`. Everytime we add an element, it represents a possible subarray - so continuously update `maxSubarray` to contain the maximum of `currentSubarray` and itself.

#### Code
```go
func maxSubarrayUsingBruteForce(arr []int) int {
	maxSubarray := math.Inf(-1)

	for i := 0; i < len(arr); i++ {
		currentSubarray := 0
		for j := i; j < len(arr); j++ {
			currentSubarray += arr[j]
			maxSubarray = math.Max(maxSubarray, float64(currentSubarray))
		}
	}
	return int(maxSubarray)
}
```
## Optimised Approach
1. Initialize 2 integer variables. Set both of them equal to the first value in the array.

- currentSubarray will keep the running count of the current subarray we are focusing on.
- maxSubarray will be our final return value. Continuously update it whenever we find a bigger subarray.
2. Iterate through the array, starting with the 2nd element (as we used the first element to initialize our variables). For each number, add it to the currentSubarray we are building. If currentSubarray becomes negative, we know it isn't worth keeping, so throw it away. Remember to update maxSubarray every time we find a new maximum.

#### Code
```go
func maxSubArray(arr []int) int {
	maxSubArray := arr[0]
	currentSubarray := arr[0]

	for i := 1; i < len(arr); i++ {
		num := arr[i]
		currentSubarray = int(math.Max(float64(num), float64(currentSubarray+num)))
		maxSubArray = int(math.Max(float64(currentSubarray), float64(maxSubArray)))
	}
	return maxSubArray
}
```