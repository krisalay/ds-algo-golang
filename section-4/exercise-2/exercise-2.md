
# Exercise 2

Given two array, write a function that lets a user know (true/false) whether these two arrays contain any common items.

For example:

| array1 | array2 | response |
|--|--|--|
| ['a', 'b', 'c', 'x'] | ['y', 'z', 'i'] | false |
| ['a', 'b', 'c', 'x'] | ['y', 'z', 'x'] | true |

  

### Jot Down Details

1. How are these collections given - in memory or array or something else? ***Response:** Array*

2. How about the elements of the array, are they chars or they can also be of type integere, float, etc? ***Response:** The elements are string type.*

3. Can the arrays fit in memory? ***Response:** For now let's say yes the arrays can fit in memory*

  

### Brute Force Approach

The most straightforward approach involves comparing every conceivable pair, achieved by implementing a nested for-loop. The outer and inner loop starts from the 0th index, systematically checking each pair to determine if value matches. While not the most efficient method, this serves as one viable solution to the problem. This approach will have quadratic time complexity.

  

### Hash Map Approach

  

Instead of inquiring, 'Is there anywhere?' we shift our perspective to, 'Have I already seen it?'. To efficiently track this, we employ a data structure suitable for lookups, such as a hash map, ensuring constant time lookup. We insert each unique value from first array into the hash map. And iterate the second array while checking the presence of that item in the hash-map. This way, we can solve this problem in linear time.

  

Time Complexity: O(m*n)

Space Complexity: O(m)

where m and n are the lengths of first and second array respectively.

  

#### Code:

```go

func  hasCommonItems(array1 []byte, array2 []byte) bool {
  hashTable := make(map[byte]bool)
  for  _, item := range array1 {
    if !hashTable[item] {
      hashTable[item] = true
    }
  } 
  for  _, item := range array2 {
    if hashTable[item] {
      return  true
    }
  }
  return  false
}
```

  

### Array is Large

What would we do differently if say we have 10 million integers in the array?

**Questions**

1. Does the array fit in the memory? ***Response:** No, the array does not fit in the memory.*

  

If our hash-map fits in memory, but input does not fit in memory, then we can sort of process it in chunks.