package algorithm

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func removeKDigits(num string, k int) string {
	var stack []byte
	for i := range num {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) != 0 || num[i] != '0' {
			stack = append(stack, num[i])
		}
	}

	if k > len(stack) {
		return "0"
	}

	t := strings.TrimLeft(string(stack[:len(stack)-k]), "0")
	if t == "" {
		return "0"
	}

	return t
}

func Test(t *testing.T) {
	require.Equal(t, "1219", removeKDigits("1432219", 3))
	require.Equal(t, "0", removeKDigits("10", 1))
}
