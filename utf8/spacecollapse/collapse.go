//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var ans strings.Builder
	ans.Grow(len(input))
	isLastRuneSpace := false
	for i := 0; i < len(input); {
		nextRune, sz := utf8.DecodeRuneInString(input[i:])
		i += sz
		if strings.ContainsRune("\t\n\r", nextRune) {
			nextRune = ' '
		}
		if !(isLastRuneSpace && nextRune == ' ') {
			ans.WriteRune(nextRune)
		}
		isLastRuneSpace = nextRune == ' '
	}
	return ans.String()
}
