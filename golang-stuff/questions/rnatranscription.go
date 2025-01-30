package questions

func ToRNA(dna string) string {
	transcribe := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	res := ""
	for _, v := range dna {
		res += string(transcribe[v])
	}
	return res
}
