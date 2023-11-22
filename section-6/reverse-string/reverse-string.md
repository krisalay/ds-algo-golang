
# Reverse String

Given a string, write a function to reverse that string.

For example:

Input String = "DS Algo using Golang"
Output String = "gnaloG gnisu oglA SD"

  

### Jot Down Details

1. Can the string fit in memory? ***Response:** For now let's say yes the arrays can fit in memory*


### Solution
The function iterates through the characters of the input string in reverse order, starting from the last character and moving towards the first character. During each iteration, it appends the current character to the backwardArray. And finally convert the array of bytes to string.

Time Complexity: O(n)

Space Complexity: O(n)


  

#### Code:

```go

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
```