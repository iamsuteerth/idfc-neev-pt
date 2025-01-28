package questions

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of DNA strands is not the same.")
	}
	hammingDistance := 0
	for i := range b {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
