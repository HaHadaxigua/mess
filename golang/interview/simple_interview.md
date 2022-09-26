# 1. Language level
## 1. How many years have you been using go, why choose go? not others

## 2. Is there anything you love/hate go very much ? And why ?

## 3. Do you have any other intersting in other language domain? Like c, java ?

## 4. If there is no slice in go, how do you implement a dynamic array?

1. allocate/free
2. gc stuff
3. language specification

## 5. Do you use reflect offen? Tell me about reflect.

## 6. Have you heard about AST?

## 7. Can you tell me something about Runtime? What/Why/How




# 2. OS level
## 1. different between process and thread
## 2.

# 3. Algorithm
## 1. Remove K digit to make the number minimum
```go
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
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right) ) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
