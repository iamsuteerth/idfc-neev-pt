package questions

import (
	"regexp"
	"strconv"
)

func Answer(s string) (int, bool) {
	match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`, s)
	if !match {
		return 0, false
	}

	re1 := regexp.MustCompile(`(plus|minus|divided|multiplied)`)
	operators := re1.FindAllString(s, -1) // Find all the occurences of the operations

	re2 := regexp.MustCompile(`-?\d+`)
	numbers := re2.FindAllString(s, -1) // Find all the occurences of the numbers

	if len(operators) != len(numbers)-1 { // For 'n' operators there are 'n - 1' operations
		return 0, false
	}
	sum, _ := strconv.Atoi(numbers[0])
	for i, o := range operators {
		n, _ := strconv.Atoi(numbers[i+1])
		switch o {
		case "plus":
			sum += n
		case "minus":
			sum -= n
		case "divided":
			sum /= n
		case "multiplied":
			sum *= n
		}
	}
	return sum, true
}
