package pangram

import "unicode"
func IsPangram(input string) bool {
	seen := make(map[rune]bool, 26)

	for _, char := range input {
		lower := unicode.ToLower(char)
		if lower >= 'a' && lower <= 'z' {
			seen[lower] = true
			if len(seen) == 26 {
				return true
			}
		}
	}

	return len(seen) == 26
}
