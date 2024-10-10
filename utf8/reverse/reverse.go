//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	var ans strings.Builder
	ans.Grow(len(input))
	for i := len(input) - 1; i >= 0; {
		r, sz := utf8.DecodeLastRuneInString(input[:i+1])
		i -= max(1, sz)
		ans.WriteRune(r)
	}
	return ans.String()
}
