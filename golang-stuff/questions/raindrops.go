package questions

import "fmt"

func Convert(number int) string {
	var res string = ""
	if number%3 == 0 {
		res += "Pling"
	}
	if number%5 == 0 {
		res += "Plang"
	}
	if number%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return fmt.Sprintf("%d", number)
	} else {
		return res
	}
}
