package questions

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i, c := range isbn {
		var digit int
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digit = int(c - '0')
		case 'X':
			if i != 9 { // Check if X is in the last position
				return false
			}
			digit = 10
		default:
			return false // Invalid character
		}
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}
