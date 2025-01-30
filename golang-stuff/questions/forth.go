package questions

import (
	"errors"
	"strconv"
	"strings"
)

// Define a type for Forth words which are operations followed by numbers
type word struct {
	numArgs int               // args to take from stack
	isValid func([]int) bool  // validate function
	run     func([]int) []int // execute
}

// Define operations
var ops = map[string]word{
	"+": {2, nil, func(a []int) []int { // 2 arguments, no validatory function and anonymous function returning sum
		return []int{a[0] + a[1]}
	}},
	"-": {2, nil, func(a []int) []int {
		return []int{a[0] - a[1]}
	}},
	"*": {2, nil, func(a []int) []int {
		return []int{a[0] * a[1]}
	}},
	"/": {2, func(a []int) bool { // Validatory function to avoid divide by zero
		return a[1] != 0
	}, func(a []int) []int {
		return []int{a[0] / a[1]}
	}},
	"swap": {2, nil, func(a []int) []int {
		return []int{a[1], a[0]}
	}},
	"over": {2, nil, func(a []int) []int {
		return []int{a[0], a[1], a[0]}
	}},
	"dup": {1, nil, func(a []int) []int {
		return []int{a[0], a[0]}
	}},
	"drop": {1, nil, func(a []int) []int {
		return []int{}
	}},
}

func Forth(lines []string) ([]int, error) {
	macros := make(map[string]string) // defined macros

	for i := range lines {
		line := strings.ToLower(lines[i])

		if line[0] == ':' { // This refers to us defining a macro
			parts := strings.SplitN(line[2:len(line)-2], " ", 2) // Extract name and definition.
			macros[parts[0]] = expandMacros(parts[1], macros)    // Store macro which can be nested
		} else {
			lines[i] = expandMacros(line, macros) // Expand any macros in the line.
		}
	}
	stack := make([]int, 0, 8)
	// Process the last line (the actual Forth program).
	for _, word := range strings.Fields(lines[len(lines)-1]) { // Split the line into words
		if op, ok := ops[word]; ok { // Check if it's a builtin op
			argsIdx := len(stack) - op.numArgs // Calculate the starting index of arguments
			if argsIdx < 0 {                   // Check for stack underflow
				return nil, errors.New("stack underflow")
			}
			if op.isValid != nil && !op.isValid(stack[argsIdx:]) { // Validate
				return nil, errors.New("invalid arguments")
			}
			stack = append(stack[:argsIdx], op.run(stack[argsIdx:])...) // Apply op
		} else if num, err := strconv.Atoi(word); err == nil { // try number parsing
			stack = append(stack, num) // Push the number into stack
		} else {
			return nil, errors.New("invalid word") // Invalid word
		}
	}

	return stack, nil // Return the final stack and any error
}

func expandMacros(line string, macros map[string]string) string {
	for name, definition := range macros {
		line = strings.ReplaceAll(line, name, definition) // replace macro names with definitions
	}
	return line
}
