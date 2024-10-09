//go:build !solution

package reverse

import "strings"

func Reverse(input string) string {
	var ans strings.Builder
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		ans.WriteRune(runes[len(runes)-i-1])
	}
	return ans.String()
}
