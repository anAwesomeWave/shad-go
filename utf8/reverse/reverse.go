//go:build !solution

package reverse

import (
	"strings"
)

func Reverse(input string) string {
	var ans strings.Builder
	ans.Grow(len(input))
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		ans.WriteRune(runes[i])
	}
	return ans.String()
}
