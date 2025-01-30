package questions

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	re := regexp.MustCompile(`\w+('\w+)?`) // Either a word or a word followed by apostrophe and word
	words := re.FindAllString(strings.ToLower(phrase), -1)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
