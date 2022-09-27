# 1. Language level

## 1. How many years have you been using go, why choose go? not others

## 2. Is there anything you love/hate go very much ? And why ?

## 3. Do you have any other interesting in other language domain? Like c, java ?

## 4. If there is no slice in go, how do you implement a dynamic array?

1. allocate/free
2. gc stuff
3. language specification

## 5. Do you use reflect often? Tell me about reflect.

## 6. Have you heard about meta-programming or AST?

## 7. Why go needs GC?

# 2. OS level

## 1. different between process and thread

# 3. Algorithm

## 1. Remove K digit to make the number minimum

```go
package algroithm

import "strings"

func removeKDigits(num string, k int) string {
	var stack []byte
	// we need to travel this array
	for i := range num {
		// check stack size and peek the top of the stack and check k size
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			// we can pop the top
			stack = stack[:len(stack)-1]
			// decrease the k
			k--
		}
		stack = append(stack, num[i])
	}

	// think about when k bigger than stack
	if k > len(stack) {
		return "0"
	}

	// think about when the number start with '0'
	t := strings.TrimLeft(string(stack[:len(stack)-k]), "0") // we should remove addititon digit
	if t == "" {
		// think about when number is an empty string
		return "0"
	}

	return t
}
```

## 2. Depth of a binary tree

```go
package algroithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 3. climb stairs

Climbing a staircase, it takes n steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

```go
package algorithm

func climbStairs(n int) int {
	dp := make([]int, n+1, n+1)
	dp[1] = 1
	if n <= 1 {
		return n
	}
	dp[2] = 2
	if n <= 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
```
