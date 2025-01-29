package questions

func All(n int, s string) []string {
	if len(s) < n {
		return nil
	}
	var result []string
	runes := []rune(s)
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, string(runes[i:n+i]))
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
