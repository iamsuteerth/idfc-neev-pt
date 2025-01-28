package questions

// Objective is to detect case insensitive anagrams but considering the fact that words are not anagrams of themselves.

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	freqMap := make(map[rune]int)
	var res []string
	for _, c := range subject {
		freqMap[unicode.ToLower(c)]++
	}
	for _, s := range candidates {
		if s == subject || strings.ToLower(subject) == strings.ToLower(s) {
			continue
		}
		testerMap := make(map[rune]int)
		for _, c := range s {
			testerMap[unicode.ToLower(c)]++
		}
		flag := true
		for k, v := range testerMap {
			if freqMap[k] != v {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, s)
		}
	}
	return res
}
