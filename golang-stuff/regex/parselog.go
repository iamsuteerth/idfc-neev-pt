package regex

import "regexp"

func IsValidLine(text string) bool {
	validPrefixRegex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return validPrefixRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)
	return re.Split(text, -1) // -1 signifies all substrings
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	// (?i) is case insensitive flag
	// .* Matches any character (.) zero or more times (*)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
	// \d+ Matches one or more digits (0-9).
}

func TagWithUserName(lines []string) []string {
	var tagged []string
	re := regexp.MustCompile(`User\s+([A-Za-z0-9]*)`) // Highlight all username strings
	// \s+ Matches one or more whitespace characters (spaces, tabs, newlines)
	// ( )  defines a capturing group
	for _, line := range lines {
		taggedLogs := re.FindStringSubmatch(line)
		// First occurrence of User...
		// The first element of the slice is the entire matched substring
		// Basically ["The quick", "The", "quick"]
		if taggedLogs != nil {
			tagged = append(tagged, "[USR] "+taggedLogs[1]+" "+line)
		} else {
			tagged = append(tagged, line)
		}
	}
	return tagged
}
