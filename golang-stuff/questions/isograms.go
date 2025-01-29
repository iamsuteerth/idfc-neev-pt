package questions

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	freq := make(map[rune]int)
	for _, c := range word {
		if c == ' ' || c == '-' {
			continue
		}
		val, exists := freq[c]
		if !exists {
			freq[c] = 1
		} else {
			if val == 1 {
				return false
			}
		}
	}
	return true
}
