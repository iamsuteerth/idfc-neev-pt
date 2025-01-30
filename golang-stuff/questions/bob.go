package questions

import (
	"strings"
	"unicode"
)

func isSilence(remark string) bool {
	return remark == ""
}

func isShouting(remark string) bool {
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0 // First letter index or -1 if none
	isUpcased := strings.ToUpper(remark) == remark
	return hasLetters && isUpcased
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func hasYelledAQuestion(remark string) bool {
	return isShouting(remark) && isQuestion(remark)
}

func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case isSilence(remark):
		return "Fine. Be that way!"
	case hasYelledAQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isShouting(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}
