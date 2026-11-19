package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, char := range word {
		// Ignore spaces, hyphens, and other non-letter characters
		if !unicode.IsLetter(char) {
			continue
		}

		lower := unicode.ToLower(char)

		// If we've already encountered this letter, it's not an isogram
		if seen[lower] {
			return false
		}

		seen[lower] = true
	}

	return true
}
