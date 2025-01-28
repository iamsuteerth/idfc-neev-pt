package questions

import (
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	// For multiple files, output is different
	var prependLineNumber, caseInsensitive, atleastOneLineMatch, exactMatch, inverseMatch bool
	for _, flag := range flags {
		switch flag {
		case "-n":
			prependLineNumber = true
		case "-i":
			caseInsensitive = true
		case "-l":
			atleastOneLineMatch = true
		case "-x":
			exactMatch = true
		case "-v":
			inverseMatch = true
		}
	}
	// Order of checking the flags :
	// caseInsensitve (requires pre processing)
	// exactMatch (string matching only but an additional parameter)
	// inverseMatch
	// atLeastOneLineMatch (since this changes the output)
	// prependLineNumber (output modification)
	for _, file := range files {
		fileContent, _ := os.ReadFile(file)
		// Going through every line which is splitted with the newline character.
		for index, line := range strings.Split(string(fileContent), "\n") {
			// Ignore Empty lines
			if line == "" {
				continue
			}
			// Before applying any flags
			processedLine, processedPattern := line, pattern
			if caseInsensitive {
				// Convert all to lower case
				processedLine, processedPattern = strings.ToLower(line), strings.ToLower(pattern)
			}
			var match bool
			if exactMatch {
				match = (processedLine == processedPattern)
			} else {
				// Contain method to search for the pattern in the line.
				match = strings.Contains(processedLine, processedPattern)
			}
			// inverseMatch and match are opposites
			if (match && !inverseMatch) || (!match && inverseMatch) {
				// If we need lines with at least one match
				var whetherPrependLineOrNot string
				if atleastOneLineMatch {
					// We only need to output name of files in case of atleastOneLineMatch flag
					result = append(result, file)
					break
				} else if prependLineNumber {
					// Add a line number as per the prepend flag
					whetherPrependLineOrNot = fmt.Sprintf("%d:%s", index+1, line)
				} else {
					whetherPrependLineOrNot = line
				}
				// If there are multiple files
				if len(files) > 1 {
					// File Name : Matched Line
					result = append(result, fmt.Sprintf("%s:%s", file, whetherPrependLineOrNot))
				} else {
					// Single File case
					result = append(result, whetherPrependLineOrNot)
				}
			}
		}
	}
	return result
}
