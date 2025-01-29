package questions

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	n := len(id)
	if n < 2 {
		return false
	}
	sum := 0
	for i := n - 1; i >= 0; i-- {
		num := id[i]
		if num < '0' || num > '9' {
			return false
		}
		digit := int(num - '0')
		if (n-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
