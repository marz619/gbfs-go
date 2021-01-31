package gbfs

import "unicode"

// ContainsRuneFunc ...
func ContainsRuneFunc(s string, f func(rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}
	return false
}

// ContainsSpaces ...
func ContainsSpaces(s string) bool {
	return ContainsRuneFunc(s, unicode.IsSpace)
}
