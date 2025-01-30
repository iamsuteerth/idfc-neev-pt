package questions

import "strings"

func IsPangram(input string) bool {
	freq := make([]int, 26)
	input = strings.ToLower(input)
	for _, c := range input {
		if c < 'a' || c > 'z' {
			continue
		}
		freq[int(c-'a')]++
	}
	for _, v := range freq {
		if v == 0 {
			return false
		}
	}
	return true
}
