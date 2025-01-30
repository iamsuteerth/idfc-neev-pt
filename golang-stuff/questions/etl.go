package questions

import "strings"

func Transform(input map[int][]string) map[string]int {
	var etl = make(map[string]int)
	for pts, letters := range input {
		for _, letter := range letters {
			etl[strings.ToLower(letter)] = pts
		}
	}
	return etl
}
